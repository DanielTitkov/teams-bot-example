package entgo

import (
	"context"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserByUsername(username string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserByTeamsID(teamsID string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.TeamsIDEQ(teamsID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) CreateUser(u *domain.User) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetUsername(u.Username).
		SetDisplayName(u.DisplayName).
		SetPasswordHash(u.PasswordHash).
		SetNillableTeamsID(u.Meta.Teams.ID).
		SetNillableTelegramID(u.Meta.Telegram.ID).
		SetNillableSlackID(u.Meta.Slack.ID).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserCount() (int, error) {
	return r.client.User.Query().Count(context.Background())
}

func entToDomainUser(user *ent.User) *domain.User {
	var email string
	if user.Email != nil {
		email = *user.Email
	}

	return &domain.User{
		ID:           user.ID,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Email:        email,
		PasswordHash: user.PasswordHash,
		Meta: domain.UserMeta{
			Teams: domain.UserMessagerData{
				ID: user.TeamsID,
			},
			Telegram: domain.UserMessagerData{
				ID: user.TelegramID,
			},
			Slack: domain.UserMessagerData{
				ID: user.SlackID,
			},
		},
	}
}
