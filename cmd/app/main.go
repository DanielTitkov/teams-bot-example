package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitkov/teams-bot-example/cmd/app/prepare"
	"github.com/DanielTitkov/teams-bot-example/internal/app"
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent"
	"github.com/DanielTitkov/teams-bot-example/internal/teams"
	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	logger.Info("starting service", "")

	cfg, err := configs.ReadConfigs("./configs/dev.yml")
	if err != nil {
		logger.Fatal("failed to load config", err)
	}
	logger.Info("loaded config", fmt.Sprintf("%+v", cfg))

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	app := app.NewApp(cfg, logger, repo)

	t := teams.NewTeams(cfg, logger)
	t.SetOnMessageHandler(app.HandleMessage)
	t.SetOnInvokeHandler(app.HandleInvoke)
	t.SetOnUpdateHandler(app.HandleUpdate)
	t.SetProactiveChannel(app.ProactiveChan)

	go func() {
		for {
			t := domain.Turn{
				Message: domain.Message{
					Text: "Proactive text",

					Attachment: string(cardJSON),
				},
				Dialog: domain.TurnDialog{
					Meta: domain.DialogMeta{
						Teams: dialogData,
					},
				},
			}

			time.Sleep(50000 * time.Second)
			app.ProactiveChan <- &t
			fmt.Println("PROACTIVE PUSHED")
		}
	}()

	go t.RunProactiveManager()

	logger.Fatal("failed to start teams service", t.Listen())

	// server := prepare.NewServer(cfg, logger, app)
	// logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}

var cardJSON = []byte(`{
	"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
	"type": "AdaptiveCard",
	"version": "1.0",
	"body": [
	  {
		"type": "TextBlock",
		"text": "This is some text",
		"size": "large"
	  },
	  {
		"type": "TextBlock",
		"text": "It doesn't wrap by default",
		"weight": "bolder"
	  },
	  {
		"type": "TextBlock",
		"text": "So set **wrap** to true if you plan on showing a paragraph of text",
		"wrap": true
	  },
	  {
		"type": "TextBlock",
		"text": "You can also use **maxLines** to prevent it from getting out of hand",
		"wrap": true,
		"maxLines": 2
	  },
	  {
		"type": "TextBlock",
		"text": "You can even draw attention to certain text with color",
		"wrap": true,
		"color": "attention"
	  }
	]
  }`)

var dialogData string = `
	{
		"activityId": "1607280724087",
		"user": {
			"id": "29:1iCnDoqN9zCOkxZgT8SkL9fpbRT_Elpv1C6t3q_9EabdrVTtdDk_cL1Vh7rDfb2Y90YEt1LUXZiY1LpmYpZp-pw",
			"name": "Daniil Titkov",
			"aadObjectId": "ffd175be-553e-4963-b6a7-84d01c6a8529"
		},
		"bot": {
			"id": "28:f7827b1a-8887-4ea3-8b1b-9465d35b71c9",
			"name": "Test2test2"
		},
		"conversation": {
			"conversationType": "personal",
			"tenantId": "6d2ee5fc-b59c-4553-aadb-e8223c654453",
			"id": "a:1AbBVHOyV8mK0t0IoVvImaZXTMsPTRcWb_BR5v_rtd8MISSK4JjJN8g8V7ChF5153Ju-mJonc5IgDdN2skzHkiuAboz8EfNvF5AX2RKrk1deJO0MjrDyLNqAsfiQKDBgd"
		},
		"channelId": "msteams",
		"serviceUrl": "https://smba.trafficmanager.net/apac/"
	}
`
