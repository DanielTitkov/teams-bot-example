package app

import (
	"errors"
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
		message.Text = buildProcessingFailedMessage(message.Err)
		return message
	}

	reply, err := a.buildReply(&message)
	if err != nil {
		message.Text = buildBuildingReplyFailedMessage(err)
		return message
	}
	// store reply

	return *reply
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
			Username:     generateUserLogin(*message.User.Teams.Username),
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

		err = a.SendTeamsProactive(&domain.Message{ // FIXME
			Text: buildUserCreatedMessage(user.DisplayName, user.Username),
			Dialog: domain.DialogMeta{
				Teams: message.Dialog.Teams,
			},
			System:    TeamsSystemCode,
			Direction: OutputMessageCode,
			Proactive: true,
		})
		if err != nil {
			a.logger.Error("failed to send user created notification", err)
		} else {
			a.logger.Info("user created notification sent", fmt.Sprint())
		}

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

func (a *App) SendTeamsProactive(m *domain.Message) error {
	if m.Dialog.Teams == "" {
		return errors.New("teams dialog reference is required to send proactive message")
	}
	a.ProactiveChan <- m // TODO: maybe add timeout
	return nil
}
