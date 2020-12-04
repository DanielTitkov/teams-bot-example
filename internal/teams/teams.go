package teams

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/teams-bot-example/internal/app"
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
}

func NewTeams(
	app *app.App,
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

func (t *Teams) ProcessMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	if t.onMessageHandler == nil {
		err := errors.New("message handler required")
		t.logger.Error("On message handler is not set", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var handler = activity.HandlerFuncs{
		OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			response := t.onMessageHandler(domain.Message{
				Text: turn.Activity.Text,
			})
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

func (t *Teams) Serve() {
	http.HandleFunc("/api/messages", t.ProcessMessage)
	port := fmt.Sprintf(":%d", t.cfg.Teams.Port)
	t.logger.Info("Listening for teams messages", "port "+port)
	http.ListenAndServe(port, nil)
}
