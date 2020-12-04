package teams

import (
	"context"
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

	activity, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		t.logger.Error("Failed to parse request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.adapter.ProcessActivity(ctx, activity, handler)
	if err != nil {
		t.logger.Error("Failed to process request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.logger.Info("Request processed successfully", "")
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
