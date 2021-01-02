package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) StoreMessage(turn mesga.Turn, dialog *domain.Dialog) error {
	m := mesgaToDomainMessage(turn.Message)
	m.System = turn.System
	message, err := a.repo.CreateMessage(&m, dialog, turn.Err) // TODO: make app method
	if err != nil {
		return err
	}
	a.logger.Info("message stored", fmt.Sprint(message))
	return nil
}
