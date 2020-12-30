package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func mesgaToDomainUser(u mesga.User) domain.User {
	return domain.User{
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Meta: domain.UserMeta{
			Teams: domain.UserMessagerData{
				ID:       u.Meta.Teams.ID,
				Username: u.Meta.Teams.Username,
			},
		},
	}
}

func mesgaToDomainDialog(mesga.Dialog) domain.Dialog {

}

func mesgaToDomainMessage(mesga.Message) domain.Message {

}

func domainToMesgaUser(mesga.User) domain.User {

}

func domainToMesgaDialog(mesga.Dialog) domain.Dialog {

}

func domainToMesgaMessage(mesga.Message) domain.Message {

}
