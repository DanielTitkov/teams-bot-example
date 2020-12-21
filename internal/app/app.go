package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
)

const (
	TeamsSystemCode   = "teams"
	InputMessageCode  = "input"
	OutputMessageCode = "output"
)

type (
	App struct {
		cfg           configs.Config
		logger        *logger.Logger
		repo          Repository
		ProactiveChan chan *domain.Turn
	}
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserByTeamsID(teamsID string) (*domain.User, error)
		GetUserCount() (int, error)

		// messages
		CreateMessage(*domain.Message, *domain.Dialog) (*domain.Message, error)

		// dialog
		CreateDialog(*domain.User, *domain.Dialog) (*domain.Dialog, error)
		GetUserDialog(*domain.User) (*domain.Dialog, error)

		// projects
		CreateProject(*domain.User, *domain.Project) (*domain.Project, error)
		GetUserProjects(*domain.User) ([]*domain.Project, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) *App {
	return &App{
		cfg:           cfg,
		logger:        logger,
		repo:          repo,
		ProactiveChan: make(chan *domain.Turn),
	}
}
