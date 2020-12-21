package app

import (
	"errors"
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

func (a *App) HandleMessage(turn domain.Turn) domain.Turn {
	turn.Message.System = TeamsSystemCode
	turn.Message.Direction = InputMessageCode
	turn.Message.Proactive = false
	user, err := a.GetOrCreateTeamsUser(turn)
	if err != nil {
		a.logger.Error("failed to get or create user", err)
		turn.Err = err
	}
	turn.User.User = user

	dialog, err := a.GetOrCreateTeamsUserDialog(turn)
	if err != nil {
		a.logger.Error("failed to get or create dialog", err)
		turn.Err = err
	}
	turn.Dialog.Dialog = dialog

	err = a.StoreMessage(turn)
	if err != nil {
		a.logger.Error("failed to store message", err)
		turn.Err = err
	}

	if turn.Err != nil {
		turn.Message.Text = buildProcessingFailedMessage(turn.Err)
		return turn
	}

	reply, err := a.buildReply(&turn)
	if err != nil {
		turn.Message.Text = buildBuildingReplyFailedMessage(err)
		return turn
	}
	// store reply

	return *reply
}

func (a *App) HandleInvoke(turn domain.Turn) domain.Turn {
	turn.Message.Text = "INVOKE"
	return turn
}

func (a *App) HandleUpdate(turn domain.Turn) domain.Turn {
	turn.Message.Text = "UPDATE"
	return turn
}

func (a *App) StoreMessage(turn domain.Turn) error {
	_, err := a.repo.CreateMessage(&turn.Message, turn.Dialog.Dialog) // TODO: make app method
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetOrCreateTeamsUserDialog(turn domain.Turn) (*domain.Dialog, error) {
	dialog, err := a.repo.GetUserDialog(turn.User.User)
	if err != nil {
		dialog, err = a.repo.CreateDialog(turn.User.User, &domain.Dialog{
			Meta: domain.DialogMeta{
				Teams: turn.Dialog.Meta.Teams,
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

func (a *App) SendTeamsProactive(t *domain.Turn) error {
	if t.Dialog.Meta.Teams == "" {
		return errors.New("teams dialog reference is required to send proactive turn")
	}
	a.ProactiveChan <- t // TODO: maybe add timeout
	return nil
}
