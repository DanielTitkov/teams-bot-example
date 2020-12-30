package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) StoreMessage(turn mesga.Turn) error {
	message, err := a.repo.CreateMessage(&turn.Message, turn.Dialog.Dialog, turn.Err) // TODO: make app method
	if err != nil {
		return err
	}
	a.logger.Info("message stored", fmt.Sprint(message))
	return nil
}
