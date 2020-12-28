package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DanielTitkov/teams-bot-example/cmd/app/prepare"
	"github.com/DanielTitkov/teams-bot-example/internal/app"
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/job"
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

	t := teams.NewTeams(logger, teams.Config{
		AppID:       cfg.Teams.AppID,
		AppPassword: cfg.Teams.AppPassword,
	})

	t.SetOnMessageHandler(app.HandleMessage)
	t.SetOnInvokeHandler(app.HandleInvoke)
	t.SetOnUpdateHandler(app.HandleUpdate)
	t.SetProactiveChannel(app.ProactiveChan)
	t.SetSentChannel(app.SentChan)

	jobs := job.NewService(cfg, logger, app)
	jobs.SendProjectNotifications() // TODO: maybe hide it inside jobs

	go t.RunProactiveManager()
	go app.ReadSentChannel()

	teamsHandler, err := t.GetHandler()
	if err != nil {
		logger.Fatal("failed to get teams handler", err)
	}

	port := fmt.Sprintf(":%d", cfg.Teams.Port)
	http.HandleFunc("/api/messages", teamsHandler)
	logger.Info("starting listening for teams messages", "port "+port)
	logger.Fatal("failed to start teams service", http.ListenAndServe(port, nil))

	// server := prepare.NewServer(cfg, logger, app)
	// logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
