package entgo

import (
	"context"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) CreateDialog(u *domain.User, d *domain.Dialog) (*domain.Dialog, error) {
	dialog, err := r.client.Dialog.
		Create().
		SetUserID(u.ID).
		SetMeta(d.Meta).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return r.entToDomainDialog(dialog), nil
}

func (r *EntgoRepository) GetUserDialog(u *domain.User) (*domain.Dialog, error) {
	dialog, err := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryDialog().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return r.entToDomainDialog(dialog), nil
}

func (r *EntgoRepository) entToDomainDialog(d *ent.Dialog) *domain.Dialog {
	return &domain.Dialog{
		ID:   d.ID,
		User: d.Edges.User.Username, // FIXME probably this won't fetch on create
		Meta: domain.DialogMeta{
			Teams:    d.Meta.Teams,
			Telegram: d.Meta.Telegram,
			Slack:    d.Meta.Slack,
		},
	}
}
