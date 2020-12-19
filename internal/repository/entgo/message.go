package entgo

import (
	"context"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateMessage(m *domain.Message, d *domain.Dialog) (*domain.Message, error) {
	message, err := r.client.Message.
		Create().
		SetDialogID(d.ID).
		SetText(m.Text).
		SetProactive(m.Proactive).
		SetSystem(m.System).
		SetDirection(m.Direction).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return r.entToDomainMessage(message), nil
}

func (r *EntgoRepository) entToDomainMessage(m *ent.Message) *domain.Message {
	return &domain.Message{
		ID:         m.ID,
		Text:       m.Text,
		Attachment: m.Attachment,
	}
}
