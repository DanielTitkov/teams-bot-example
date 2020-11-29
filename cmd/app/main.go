package main

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
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

	// app := app.NewApp(cfg, logger, repo)

	// server := prepare.NewServer(cfg, logger, app)
	// logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
