package app

import "github.com/DanielTitkov/teams-bot-example/internal/domain"

func (a *App) HandleMessage(message domain.Message) domain.Message {
	message.Text = "Echo: " + message.Text
	return message
}

func (a *App) HandleInvoke(message domain.Message) domain.Message {
	message.Text = "INVOKE"
	return message
}

func (a *App) HandleUpdate(message domain.Message) domain.Message {
	message.Text = "UPDATE"
	return message
}
