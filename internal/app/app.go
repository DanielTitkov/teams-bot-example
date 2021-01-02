package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
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
		ProactiveChan chan *mesga.Turn
		SentChan      chan *mesga.Turn
	}
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserByTeamsID(teamsID string) (*domain.User, error)
		GetUserCount() (int, error)

		// messages
		CreateMessage(*domain.Message, *domain.Dialog, error) (*domain.Message, error)

		// dialog
		CreateDialog(*domain.User, *domain.Dialog) (*domain.Dialog, error)
		GetUserDialog(*domain.User) (*domain.Dialog, error)

		// projects
		CreateProject(*domain.User, *domain.Project) (*domain.Project, error)
		GetUserProjects(*domain.User) ([]*domain.Project, error)
		GetRandomProjectByUser() ([]*domain.Project, error) // for sample notifications
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
		ProactiveChan: make(chan *mesga.Turn, 100),
		SentChan:      make(chan *mesga.Turn, 100),
	}
}
