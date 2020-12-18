package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
)

type (
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
	}
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserByTeamsID(teamsID string) (*domain.User, error)
		GetUserCount() (int, error)

		// messages
		// CreateDialog(*domain.User, *domain.Dialog) (*domain.Dialog, error)
		// GetUserDialog(*domain.User) (*domain.Dialog, error)
		// CreateMessage(*domain.Message) (*domain.Message, error)

		// projects
		// CreateProject()
		// GetProjects()
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}
