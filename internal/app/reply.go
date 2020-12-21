package app

// FIXME maybe reply is not good file name

import (
	"errors"
	"regexp"
	"strings"
	"time"

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
	reply := &domain.Turn{
		Message: domain.Message{
			Direction: OutputMessageCode,
			System:    turn.Message.System,
			Proactive: false,
		},
	}

	tokens := strings.Split(turn.Message.Text, " ")
	if len(tokens) < 4 {
		return nil, errors.New(buildCreateProjectFailedMessage(errors.New("not enough data to create project, check input")))
	}
	projectTitle := tokens[2]
	projectDueDate, err := time.Parse(defaultDateTimeLayout, strings.TrimSpace(tokens[3]))
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}
	project, err := a.CreateProject(turn, &domain.Project{
		User:    turn.User.User.Username,
		Title:   projectTitle,
		DueDate: projectDueDate,
	})
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}

	reply.Message.Text = buildCreateProjectSuccessMessage(projectTitle, projectDueDate, project.ID)

	return reply, nil
}

func matchWithRegexp(text string, reg *regexp.Regexp) bool {
	return reg.Match([]byte(text))
}
