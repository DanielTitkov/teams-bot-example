package app

import (
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) CreateProject(u *domain.User, p *domain.Project) (*domain.Project, error) {
	project, err := a.repo.CreateProject(u, p)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (a *App) GetUserProjects(u *domain.User) ([]*domain.Project, error) { // TODO: think of signature
	projects, err := a.repo.GetUserProjects(u)
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
		user, err := a.repo.GetUserByUsername(p.User)
		if err != nil {
			return err
		}

		dialog, err := a.repo.GetUserDialog(&domain.User{Username: p.User})
		if err != nil {
			return err
		}

		err = a.SendTeamsProactive(&mesga.Turn{
			System: mesga.TeamsCode,
			User: &mesga.User{
				Teams: mesga.UserMessagerData{
					ID: user.Meta.Teams.ID,
				},
			},
			Dialog: &mesga.Dialog{
				Teams: dialog.Meta.Teams,
			},
			Message: mesga.Message{
				Text: buildProjectNotificationText(
					p.Title,
					p.ID,
					minutesFromNow(p.DueDate),
				),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func minutesFromNow(t time.Time) int64 {
	now := time.Now()
	diff := now.Sub(t)
	return int64(diff.Minutes()) * -1
}
