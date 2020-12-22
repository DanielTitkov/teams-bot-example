package entgo

import (
	"context"
	"math/rand"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) CreateProject(u *domain.User, p *domain.Project) (*domain.Project, error) {
	project, err := r.client.Project.
		Create().
		SetUserID(u.ID).
		SetTitle(p.Title).
		SetDueDate(p.DueDate).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return r.entToDomainProject(u, project), nil
}

func (r *EntgoRepository) GetUserProjects(u *domain.User) ([]*domain.Project, error) {
	projects, err := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryProjects().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	var res []*domain.Project
	for _, p := range projects {
		res = append(res, r.entToDomainProject(u, p))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomProjectByUser() ([]*domain.Project, error) {
	users, err := r.client.User.
		Query().
		WithProjects().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var res []*domain.Project
	for _, u := range users {
		projects := u.Edges.Projects
		rnd := rand.New(rand.NewSource(time.Now().Unix()))
		project := projects[rnd.Intn(len(projects))]
		res = append(res, r.entToDomainProject(&domain.User{Username: u.Username}, project))
	}
	return res, nil
}

func (r *EntgoRepository) entToDomainProject(u *domain.User, p *ent.Project) *domain.Project {
	return &domain.Project{
		ID:      p.ID,
		User:    u.Username,
		Title:   p.Title,
		DueDate: p.DueDate,
	}
}
