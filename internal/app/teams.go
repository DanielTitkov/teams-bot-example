package app

import (
	"errors"
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) HandleMessage(turn mesga.Turn) (updatedTurn mesga.Turn) {
	updatedTurn = turn
	defer func() {
		var err error
		if r := recover(); r != nil { // capture eventual panic in business logic
			err = fmt.Errorf("panic occured during message processing: %s", r)
			a.logger.Error("teams handle message paniced", err)
			updatedTurn.Err = err
		}
		updatedTurn.Message.Direction = mesga.InputCode
		updatedTurn.Message.Proactive = false
		updatedTurn.System = mesga.TeamsCode
		if updatedTurn.Err != nil {
			updatedTurn.Message.Text = buildBuildingReplyFailedMessage(updatedTurn.Err)
		}
	}()

	user, dialog, err := a.GetUserAndDialog(updatedTurn)
	if err != nil {
		updatedTurn.Err = err
		return updatedTurn
	}

	err = a.StoreMessage(updatedTurn, dialog)
	if err != nil {
		a.logger.Error("failed to store message", err)
		updatedTurn.Err = err
		return updatedTurn
	}

	reply, err := a.buildReply(&updatedTurn, user, dialog)
	if err != nil {
		updatedTurn.Message.Text = buildBuildingReplyFailedMessage(err)
		return updatedTurn
	}
	updatedTurn = *reply
	return updatedTurn
}

func (a *App) HandleInvoke(turn mesga.Turn) mesga.Turn {
	turn.Message.Text = "INVOKE"
	return turn
}

func (a *App) HandleUpdate(turn mesga.Turn) mesga.Turn {
	turn.Message.Text = "UPDATE"
	return turn
}

func (a *App) SendTeamsProactive(t *mesga.Turn) error {
	if t.User == nil || *t.User.Teams.ID == "" {
		return errors.New("teams user id is required to send proactive turn")
	} // TODO: maybe fetch user by dialog if user id is not present
	if t.Dialog == nil || t.Dialog.Teams == "" {
		return errors.New("teams dialog reference is required to send proactive turn")
	}
	a.ProactiveChan <- t // TODO: maybe add timeout
	return nil
}

func (a *App) ReadSentChannel() {
	for turn := range a.SentChan {
		_, dialog, err := a.GetUserAndDialog(*turn)
		if err != nil {
			a.logger.Error("failed to get user dialog", err)
		}
		err = a.StoreMessage(*turn, dialog)
		if err != nil {
			a.logger.Error("failed to store message", err)
		}
	}
	a.logger.Warn("sent channel is closed", "")
}

func (a *App) GetUserAndDialog(turn mesga.Turn) (*domain.User, *domain.Dialog, error) {
	user, err := a.GetOrCreateTeamsUser(turn)
	if err != nil {
		a.logger.Error("failed to get or create user", err)
		return nil, nil, err
	}

	dialog, err := a.GetOrCreateTeamsUserDialog(turn, user)
	if err != nil {
		a.logger.Error("failed to get or create dialog", err)
		return nil, nil, err
	}

	return user, dialog, nil
}
