package app

import (
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) CreateProject(t *mesga.Turn, p *domain.Project) (*domain.Project, error) {
	project, err := a.repo.CreateProject(t.User.User, p)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (a *App) GetUserProjects(t *mesga.Turn) ([]*domain.Project, error) { // TODO: think of signature
	projects, err := a.repo.GetUserProjects(t.User.User)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// SendProjectNotifications implementation sucks.
// It's only for testing.
func (a *App) SendProjectNotifications() error {
	projects, err := a.repo.GetRandomProjectByUser()
	if err != nil {
		return err
	}
	for _, p := range projects {
		dialog, err := a.repo.GetUserDialog(&domain.User{Username: p.User})
		if err != nil {
			return err
		}

		err = a.SendTeamsProactive(&mesga.Turn{
			Dialog: mesga.TurnDialog{
				Dialog: dialog,
				Meta:   dialog.Meta,
			},
			Message: domain.Message{
				Text: buildProjectNotificationText(
					p.Title,
					p.ID,
					minutesFormNow(p.DueDate),
				),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func minutesFormNow(t time.Time) int64 {
	now := time.Now()
	diff := now.Sub(t)
	return int64(diff.Minutes()) * -1
}
