package app

// FIXME maybe reply is not good file name

import (
	"regexp"
	"strings"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
)

var replyMapping = map[string]func(*domain.Turn) (*domain.Turn, error){}

func (a *App) buildReply(turn *domain.Turn) (*domain.Turn, error) {
	turn.Message.Direction = OutputMessageCode
	text := turn.Message.Text

	switch {
	case matchWithRegexp(text, createProjectRequest):
		return a.createProjectReply(turn)
	default:
		return a.defaultReply(turn)
	}
}

func (a *App) defaultReply(turn *domain.Turn) (*domain.Turn, error) {
	return &domain.Turn{
		Message: domain.Message{
			Text:      defaultReplyText,
			Direction: OutputMessageCode,
			System:    turn.Message.System,
			Proactive: false,
		},
	}, nil
}

func (a *App) createProjectReply(turn *domain.Turn) (*domain.Turn, error) {
	var reply string
	tokens := strings.Split(turn.Message.Text, " ")
	if len(tokens) < 4 {
		reply = buildCreateProjectFailedMessage()
	}
	reply = buildCreateProjectSuccessMessage(tokens[2], tokens[3])
	return &domain.Turn{
		Message: domain.Message{
			Text:      reply,
			Direction: OutputMessageCode,
			System:    turn.Message.System,
			Proactive: false,
		},
	}, nil
}

func matchWithRegexp(text string, reg *regexp.Regexp) bool {
	return reg.Match([]byte(text))
}
