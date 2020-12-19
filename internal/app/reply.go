package app

import "github.com/DanielTitkov/teams-bot-example/internal/domain"

var replyMapping = map[string]func(*domain.Message) (*domain.Message, error){}

func (a *App) buildReply(message *domain.Message) (*domain.Message, error) {
	message.Direction = OutputMessageCode

	builderFn, ok := replyMapping[message.Text]
	if !ok {
		return a.defaultReply(message)
	}

	return builderFn(message)
}

func (a *App) defaultReply(message *domain.Message) (*domain.Message, error) {
	return &domain.Message{
		Text:      defaultReplyText,
		Direction: OutputMessageCode,
		System:    message.System,
		Proactive: false,
	}, nil
}
