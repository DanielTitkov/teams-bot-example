package job

import (
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/app"
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
)

type Service struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewService(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Service {
	return &Service{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}

func (s *Service) SendProjectNotifications() {
	period := s.cfg.Job.ProjectNotificationPeriod
	if period == 0 {
		return
	}
	go func() {
		for {
			time.Sleep(time.Duration(period) * time.Second)
			if err := s.app.SendProjectNotifications(); err != nil {
				s.logger.Error("failed to send project notifications", err)
			} else {
				s.logger.Info("sending project notification", "")
			}
		}
	}()
}
