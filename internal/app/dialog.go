package app

import (
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

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
	} else {
		a.logger.Info("dialog fetched", fmt.Sprint(dialog))
	}
	return dialog, nil
}
