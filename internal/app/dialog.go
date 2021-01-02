package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) GetOrCreateTeamsUserDialog(turn mesga.Turn, user *domain.User) (*domain.Dialog, error) {
	dialog, err := a.repo.GetUserDialog(user)
	if err != nil {
		dialog, err = a.repo.CreateDialog(user, &domain.Dialog{
			Meta: domain.DialogMeta{
				Teams: turn.Dialog.Teams,
			},
		})
		if err != nil {
			return nil, err
		}
		a.logger.Info("dialog created", fmt.Sprint(dialog))
	} else {
		a.logger.Info("dialog fetched", fmt.Sprint(dialog))
	}
	return dialog, nil
}
