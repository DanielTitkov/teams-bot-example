package app

import (
	"errors"
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

func (a *App) HandleMessage(turn domain.Turn) (reply domain.Turn) {
	reply = turn
	defer func() { // capture eventual panic in business logic
		if r := recover(); r != nil {
			err := fmt.Errorf("panic occured during message processing: %s", r)
			a.logger.Error("teams handle message paniced", err)
			reply.Err = err
			reply.Message.Text = buildBuildingReplyFailedMessage(reply.Err)
		}
	}()

	reply.Message.System = TeamsSystemCode
	reply.Message.Direction = InputMessageCode
	reply.Message.Proactive = false

	user, err := a.GetOrCreateTeamsUser(turn)
	if err != nil {
		a.logger.Error("failed to get or create user", err)
		reply.Err = err
	} else {
		reply.User.User = user
	}

	dialog, err := a.GetOrCreateTeamsUserDialog(turn)
	if err != nil {
		a.logger.Error("failed to get or create dialog", err)
		reply.Err = err
	} else {
		reply.Dialog.Dialog = dialog
	}

	err = a.StoreMessage(turn)
	if err != nil {
		a.logger.Error("failed to store message", err)
		reply.Err = err
	}

	if turn.Err != nil {
		reply.Message.Text = buildProcessingFailedMessage(turn.Err)
		return reply
	}

	builtReply, err := a.buildReply(&turn)
	if err != nil {
		reply.Message.Text = buildBuildingReplyFailedMessage(err)
		return reply
	}

	return *builtReply
}

func (a *App) HandleInvoke(turn domain.Turn) domain.Turn {
	turn.Message.Text = "INVOKE"
	return turn
}

func (a *App) HandleUpdate(turn domain.Turn) domain.Turn {
	turn.Message.Text = "UPDATE"
	return turn
}

func (a *App) SendTeamsProactive(t *domain.Turn) error {
	if t.Dialog.Meta.Teams == "" {
		return errors.New("teams dialog reference is required to send proactive turn")
	}
	t.Message.Proactive = true
	t.Message.Direction = OutputMessageCode
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
