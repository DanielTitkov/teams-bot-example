package teams

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

var customHandler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		return turn.SendActivity(activity.MsgOptionText("Echo: " + turn.Activity.Text))
	},
}

// HTTPHandler handles the HTTP requests from then connector service
type Teams struct {
	adapter core.Adapter
	cfg     configs.Config
	logger  *logger.Logger
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

func (t *Teams) ProcessMessage(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()
	activity, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		t.logger.Error("Failed to parse request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.adapter.ProcessActivity(ctx, activity, customHandler)
	if err != nil {
		fmt.Println("Failed to process request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.logger.Info("Request processed successfully", "")
}

func (t *Teams) Serve() {
	http.HandleFunc("/api/messages", t.ProcessMessage)
	t.logger.Info("Listening for teams messages on port:3978...", "")
	http.ListenAndServe(":3978", nil)
}
