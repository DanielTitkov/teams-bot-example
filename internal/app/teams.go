package app

import "github.com/DanielTitkov/teams-bot-example/internal/domain"

func (a *App) HandleMessage(message domain.Message) domain.Message {
	message.Text = "Echo: " + message.Text
	return message
}
