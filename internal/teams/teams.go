package teams

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

// Teams handles the HTTP requests from then connector service
type Teams struct {
	adapter          core.Adapter
	cfg              configs.Config
	logger           *logger.Logger
	onMessageHandler func(domain.Message) domain.Message
	onInvokeHandler  func(domain.Message) domain.Message
	onUpdateHandler  func(domain.Message) domain.Message
	proactiveChan    <-chan domain.Message
}

func NewTeams(
	cfg configs.Config,
	logger *logger.Logger,
) *Teams {
	setting := core.AdapterSetting{
		AppID:       cfg.Teams.AppID,
		AppPassword: cfg.Teams.AppPassword,
	}

	adapter, err := core.NewBotAdapter(setting)
	if err != nil {
		log.Fatal("Error creating adapter: ", err)
	}

	return &Teams{
		adapter: adapter,
		cfg:     cfg,
		logger:  logger,
	}
}

func (t *Teams) SetOnMessageHandler(handler func(domain.Message) domain.Message) {
	t.onMessageHandler = handler
}

func (t *Teams) SetOnInvokeHandler(handler func(domain.Message) domain.Message) {
	t.onInvokeHandler = handler
}

func (t *Teams) SetOnUpdateHandler(handler func(domain.Message) domain.Message) {
	t.onUpdateHandler = handler
}

func (t *Teams) SetProactiveChannel(ch chan domain.Message) {
	t.proactiveChan = ch
}

func (t *Teams) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	var handler = activity.HandlerFuncs{
		OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			response := t.onMessageHandler(domain.Message{
				Text: turn.Activity.Text,
			})
			return turn.SendActivity(activity.MsgOptionText(response.Text))
		},
		OnInvokeFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			response := t.onInvokeHandler(domain.Message{})
			return turn.SendActivity(activity.MsgOptionText(response.Text))
		},
		OnConversationUpdateFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			response := t.onUpdateHandler(domain.Message{})
			return turn.SendActivity(activity.MsgOptionText(response.Text))
		},
	}

	act, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		t.logger.Error("Failed to parse request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set conversation reference
	var conversationRef schema.ConversationReference
	conversationRef = activity.GetCoversationReference(act)
	fmt.Printf("REF %+v\n", conversationRef) // TODO store somewhere ref

	err = t.adapter.ProcessActivity(ctx, act, handler)
	if err != nil {
		t.logger.Error("Failed to process request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.logger.Info("Request processed successfully", "")

	// Send proactive message
}

func (t *Teams) Listen() error {
	if t.onMessageHandler == nil {
		err := errors.New("message handler required")
		return err
	}

	if t.onInvokeHandler == nil {
		err := errors.New("invoke handler required")
		return err
	}

	if t.onUpdateHandler == nil {
		err := errors.New("update handler required")
		return err
	}

	http.HandleFunc("/api/messages", t.processMessage)
	port := fmt.Sprintf(":%d", t.cfg.Teams.Port)
	t.logger.Info("Listening for teams messages", "port "+port)
	http.ListenAndServe(port, nil)
	return nil
}

func (t *Teams) RunProactiveManager() {
	for {
		select {
		case message := <-t.proactiveChan:
			t.sendMessage(message)
		}
	}
}

func (t *Teams) sendMessage(message domain.Message) {
	var ref schema.ConversationReference
	err := json.Unmarshal([]byte(message.DialogData), &ref)
	if err != nil {
		t.logger.Error("Failed to unmarshal conversation reference", err)
		return
	}

	var proactiveHandler = activity.HandlerFuncs{
		OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			var obj map[string]interface{}
			err := json.Unmarshal([]byte(message.Attachment), &obj)
			if err != nil {
				return schema.Activity{}, err
			}
			attachments := []schema.Attachment{
				{
					ContentType: "application/vnd.microsoft.card.adaptive",
					Content:     obj,
				},
			}
			return turn.SendActivity(activity.MsgOptionText("Sample attachment"), activity.MsgOptionAttachments(attachments))
		},
	}

	err = t.adapter.ProactiveMessage(context.TODO(), ref, proactiveHandler)
	if err != nil {
		t.logger.Error("Failed to send proactive message.", err)
		return
	}
	t.logger.Info("Proactive message sent successfully.", fmt.Sprintf("%+v", message))
}