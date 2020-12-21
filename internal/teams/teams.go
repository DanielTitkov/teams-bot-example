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
	onMessageHandler func(domain.Turn) domain.Turn
	onInvokeHandler  func(domain.Turn) domain.Turn
	onUpdateHandler  func(domain.Turn) domain.Turn
	proactiveChan    <-chan *domain.Turn
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
		log.Fatal("error creating adapter: ", err)
	}

	return &Teams{
		adapter: adapter,
		cfg:     cfg,
		logger:  logger,
	}
}

func (t *Teams) SetOnMessageHandler(handler func(domain.Turn) domain.Turn) {
	t.onMessageHandler = handler
}

func (t *Teams) SetOnInvokeHandler(handler func(domain.Turn) domain.Turn) {
	t.onInvokeHandler = handler
}

func (t *Teams) SetOnUpdateHandler(handler func(domain.Turn) domain.Turn) {
	t.onUpdateHandler = handler
}

func (t *Teams) SetProactiveChannel(ch chan *domain.Turn) {
	t.proactiveChan = ch
}

func (t *Teams) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	var handler = activity.HandlerFuncs{
		OnMessageFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onMessageHandler(t.activityToTurn(turnCtx))
			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text))
		},
		OnInvokeFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onInvokeHandler(t.activityToTurn(turnCtx))
			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text))
		},
		OnConversationUpdateFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onUpdateHandler(t.activityToTurn(turnCtx))
			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text))
		},
	}

	act, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		t.logger.Error("failed to parse request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.adapter.ProcessActivity(ctx, act, handler)
	if err != nil {
		t.logger.Error("failed to process request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.logger.Info("request processed successfully", "")
}

func (t *Teams) Listen() error {
	if t.onMessageHandler == nil {
		return errors.New("message handler required")
	}

	if t.onInvokeHandler == nil {
		return errors.New("invoke handler required")
	}

	if t.onUpdateHandler == nil {
		return errors.New("update handler required")
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

func (t *Teams) sendMessage(turn *domain.Turn) {
	var ref schema.ConversationReference
	err := json.Unmarshal([]byte(turn.Dialog.Meta.Teams), &ref)
	if err != nil {
		t.logger.Error("failed to unmarshal conversation reference", err)
		return
	}

	var proactiveHandler = activity.HandlerFuncs{
		OnMessageFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			// var obj map[string]interface{}
			// err := json.Unmarshal([]byte(message.Attachment), &obj)
			// if err != nil {
			// 	return schema.Activity{}, err
			// }
			// attachments := []schema.Attachment{
			// 	{
			// 		ContentType: "application/vnd.microsoft.card.adaptive",
			// 		Content:     obj,
			// 	},
			// }
			return turnCtx.SendActivity(
				activity.MsgOptionText(turn.Message.Text),
				// activity.MsgOptionAttachments(attachments),
			)
		},
	}

	err = t.adapter.ProactiveMessage(context.Background(), ref, proactiveHandler)
	if err != nil {
		t.logger.Error("Failed to send proactive message.", err)
		return
	}
	t.logger.Info("Proactive message sent successfully.", fmt.Sprintf("%+v", turn.Message))
}

func (t *Teams) activityToTurn(turnCtx *activity.TurnContext) domain.Turn {
	conversationRef := activity.GetCoversationReference(turnCtx.Activity)
	jsonRef, err := json.Marshal(conversationRef)
	return domain.Turn{
		Message: domain.Message{
			Text: turnCtx.Activity.Text,
		},
		Dialog: domain.TurnDialog{
			Meta: domain.DialogMeta{
				Teams: string(jsonRef),
			},
		},
		User: domain.TurnUser{
			Meta: domain.UserMeta{
				Teams: domain.UserMessagerData{
					ID:       &conversationRef.User.ID,
					Username: &conversationRef.User.Name,
				},
			},
		},
		Err: err,
	}
}
