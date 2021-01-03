package app

import (
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func mesgaToDomainUser(u mesga.User) domain.User {
	return domain.User{
		Meta: domain.UserMeta{
			Teams: domain.UserMessagerData{
				ID:       u.Teams.ID,
				Username: u.Teams.Username,
			},
		},
	}
}

func mesgaToDomainDialog(d mesga.Dialog) domain.Dialog {
	return domain.Dialog{
		Meta: domain.DialogMeta{
			Teams: d.Teams,
		},
	}
}

func mesgaToDomainMessage(m mesga.Message) domain.Message {
	return domain.Message{
		Text:         m.Text,
		PayloadType:  m.Payload.Type,
		PayloadValue: m.Payload.Value,
		Attachment:   m.Attachment,
		Direction:    m.Direction,
		Proactive:    m.Proactive,
	}
}

func domainToMesgaUser(u domain.User) mesga.User {
	return mesga.User{
		Teams: mesga.UserMessagerData{
			ID: u.Meta.Teams.ID,
		},
	}
}

func domainToMesgaDialog(d domain.Dialog) mesga.Dialog {
	return mesga.Dialog{
		Teams: d.Meta.Teams,
	}
}

func domainToMesgaMessage(m domain.Message) mesga.Message {
	return mesga.Message{
		Text:       m.Text,
		Attachment: m.Attachment,
		Direction:  m.Direction,
		Proactive:  m.Proactive,
		Payload: mesga.MessagePayload{
			Type:  m.PayloadType,
			Value: m.PayloadValue,
		},
	}
}
