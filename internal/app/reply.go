package app

// FIXME maybe reply is not good file name

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func (a *App) buildReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	turn.Message.Direction = OutputMessageCode
	text := turn.Message.Text

	switch {
	case matchWithRegexp(text, createProjectRequest):
		return a.createProjectReply(turn, user, dialog)
	case matchWithRegexp(text, listProjiectsRequest):
		return a.listProjectsReply(turn, user, dialog)
	default:
		return a.defaultReply(turn, user, dialog)
	}
}

func (a *App) defaultReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)
	reply.Message.Text = defaultReplyText
	reply.Message.Attachment = introCardJSON
	return reply, nil
}

func (a *App) createProjectReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)

	tokens := strings.Split(turn.Message.Text, " ")
	if len(tokens) < 4 {
		return nil, errors.New(buildCreateProjectFailedMessage(errors.New("not enough data to create project, check input")))
	}
	projectTitle := tokens[2]
	projectDueDate, err := time.Parse(defaultDateTimeLayout, strings.TrimSpace(tokens[3]))
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}
	project, err := a.CreateProject(user, &domain.Project{
		User:    user.Username,
		Title:   projectTitle,
		DueDate: projectDueDate,
	})
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}

	reply.Message.Text = buildCreateProjectSuccessMessage(projectTitle, projectDueDate, project.ID)

	return reply, nil
}

func (a *App) listProjectsReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)
	projects, err := a.GetUserProjects(user)
	if err != nil {
		return nil, errors.New(buildListProjectsFailedMessage(err))
	}
	if len(projects) < 1 {
		reply.Message.Text = buildNoProjectsText()
	} else {
		replyText := buildListProjectsHeader()
		for i, p := range projects {
			replyText += buildListProjectsLine(i+1, p.ID, p.Title, p.DueDate)
		}
		reply.Message.Text = replyText
	}

	return reply, nil
}

func matchWithRegexp(text string, reg *regexp.Regexp) bool {
	return reg.Match([]byte(text))
}

func makeOutputTurn(turn *mesga.Turn) *mesga.Turn {
	return &mesga.Turn{
		System: mesga.TeamsCode,
		Message: mesga.Message{
			Direction: mesga.OutputCode,
			Proactive: false,
		},
		Dialog: turn.Dialog,
		User:   turn.User,
	}
}
