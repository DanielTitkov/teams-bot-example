package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

func (a *App) CreateProject(t *domain.Turn, p *domain.Project) (*domain.Project, error) {
	project, err := a.repo.CreateProject(t.User.User, p)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (a *App) GetUserProjects(t *domain.Turn) ([]*domain.Project, error) { // TODO: think of signature
	projects, err := a.repo.GetUserProjects(t.User.User)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
