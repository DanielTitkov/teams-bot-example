package app

// FIXME maybe reply is not good file name

import (
	"regexp"
	"strings"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

var replyMapping = map[string]func(*domain.Message) (*domain.Message, error){}

func (a *App) buildReply(message *domain.Message) (*domain.Message, error) {
	message.Direction = OutputMessageCode
	text := message.Text

	switch {
	case matchWithRegexp(text, createProjectRequest):
		return a.createProjectReply(message)
	default:
		return a.defaultReply(message)
	}
}

func (a *App) defaultReply(message *domain.Message) (*domain.Message, error) {
	return &domain.Message{
		Text:      defaultReplyText,
		Direction: OutputMessageCode,
		System:    message.System,
		Proactive: false,
	}, nil
}

func (a *App) createProjectReply(message *domain.Message) (*domain.Message, error) {
	var reply string
	tokens := strings.Split(message.Text, " ")
	if len(tokens) < 4 {
		reply = buildCreateProjectFailedMessage()
	}
	reply = buildCreateProjectSuccessMessage(tokens[2], tokens[3])
	return &domain.Message{
		Text:      reply,
		Direction: OutputMessageCode,
		System:    message.System,
		Proactive: false,
	}, nil
}

func matchWithRegexp(text string, reg *regexp.Regexp) bool {
	return reg.Match([]byte(text))
}
