package app

import (
	"errors"
	"fmt"

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
		updatedTurn.Message.System = TeamsSystemCode
		updatedTurn.Message.Direction = InputMessageCode
		updatedTurn.Message.Proactive = false
		if updatedTurn.Err != nil {
			updatedTurn.Message.Text = buildBuildingReplyFailedMessage(updatedTurn.Err)
		}
	}()

	user, err := a.GetOrCreateTeamsUser(turn)
	if err != nil {
		a.logger.Error("failed to get or create user", err)
		updatedTurn.Err = err
		return updatedTurn
	} else {
		updatedTurn.User.User = user
	}

	dialog, err := a.GetOrCreateTeamsUserDialog(updatedTurn)
	if err != nil {
		a.logger.Error("failed to get or create dialog", err)
		updatedTurn.Err = err
		return updatedTurn
	} else {
		updatedTurn.Dialog.Dialog = dialog
	}

	err = a.StoreMessage(updatedTurn)
	if err != nil {
		a.logger.Error("failed to store message", err)
		updatedTurn.Err = err
		return updatedTurn
	}

	reply, err := a.buildReply(&updatedTurn)
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
	if t.Dialog.Meta.Teams == "" {
		return errors.New("teams dialog reference is required to send proactive turn")
	}
	t.Message.Proactive = true
	t.Message.Direction = OutputMessageCode
	t.System = mesga.TeamsCode
	a.ProactiveChan <- t // TODO: maybe add timeout
	return nil
}

func (a *App) ReadSentChannel() {
	for {
		select {
		case turn := <-a.SentChan:
			err := a.StoreMessage(*turn)
			if err != nil {
				a.logger.Error("failed to store message", err)
			}
		}
	}
}
