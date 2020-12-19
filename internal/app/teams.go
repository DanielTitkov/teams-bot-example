package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) HandleMessage(message domain.Message) domain.Message {
	message.System = TeamsSystemCode
	message.Direction = InputMessageCode
	message.Proactive = false
	user, err := a.GetOrCreateTeamsUser(message)
	if err != nil {
		a.logger.Error("failed to get or create user", err)
		message.Err = err
	}

	dialog, err := a.GetOrCreateTeamsUserDialog(message, user)
	if err != nil {
		a.logger.Error("failed to get or create dialog", err)
		message.Err = err
	}

	_, err = a.repo.CreateMessage(&message, dialog)
	if err != nil {
		a.logger.Error("failed to create message", err)
		message.Err = err
	}

	if message.Err != nil {
		message.Text = fmt.Sprintf("error occured during processing message: %s", message.Err.Error())
		return message
	}

	// call message replier
	// store reply

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

func (a *App) GetOrCreateTeamsUser(message domain.Message) (*domain.User, error) {
	user, err := a.repo.GetUserByTeamsID(*message.User.Teams.ID)
	if err != nil {
		password := "123" // FIXME
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		a.logger.Info("user created", fmt.Sprint(user))
		// TODO: add proacitve - user created
	} else {
		a.logger.Info("user fetched", fmt.Sprint(user))
	}
	return user, nil
}

func (a *App) GetOrCreateTeamsUserDialog(message domain.Message, user *domain.User) (*domain.Dialog, error) {
	dialog, err := a.repo.GetUserDialog(user)
	if err != nil {
		dialog, err = a.repo.CreateDialog(user, &domain.Dialog{
			Meta: domain.DialogMeta{
				Teams: message.Dialog.Teams,
			},
		})
		if err != nil {
			return nil, err
		}
		a.logger.Info("dialog created", fmt.Sprint(dialog))
		// TODO: add proacitve - dialog created
	} else {
		a.logger.Info("dialog fetched", fmt.Sprint(dialog))
	}
	return dialog, nil
}
