package main

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
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

	t := teams.NewTeams(cfg, logger)
	t.Serve()

	// app := app.NewApp(cfg, logger, repo)

	// server := prepare.NewServer(cfg, logger, app)
	// logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
