package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) HandleMessage(message domain.Message) domain.Message {
	user, err := a.repo.GetUserByTeamsID(*message.User.Teams.ID)
	if err != nil {
		password := "123" // FIXME
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			message.Err = err
		}

		user, err = a.repo.CreateUser(&domain.User{
			DisplayName:  *message.User.Teams.Username,
			Username:     *message.User.Teams.Username, // TODO: generate unique name
			PasswordHash: string(hash),
			Email:        "sample@email.com", // FIXME
			Meta: domain.UserMeta{
				Teams: domain.UserMessagerData{
					ID: message.User.Teams.ID,
				},
			},
		})
		if err != nil {
			message.Err = err
			a.logger.Error("failed to create user", err)
		} else {
			a.logger.Info("user created", fmt.Sprint(user))
			// TODO: add proacitve - user created
		}
	} else {
		a.logger.Info("user fetched", fmt.Sprint(user))
	}

	message.Text = "Echo for user " + user.DisplayName + ": " + message.Text
	return message
}

func (a *App) HandleInvoke(message domain.Message) domain.Message {
	message.Text = "INVOKE"
	return message
}

func (a *App) HandleUpdate(message domain.Message) domain.Message {
	message.Text = "UPDATE"
	return message
}
