package app

// FIXME maybe reply is not good file name

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

const (
	createProjectAction     = "createProject"
	showProjectsAction      = "showProjects"
	initCreateProjectAction = "initCreateProject"
)

func (a *App) buildReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	turn.Message.Direction = OutputMessageCode
	if turn.Message.Payload.Value != "" && turn.Message.Payload.Value != "null" {
		return a.buildPayloadReply(turn, user, dialog)
	}
	return a.buildTextReply(turn, user, dialog)
}

func (a *App) buildTextReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	text := turn.Message.Text

	switch {
	case matchWithRegexp(text, createProjectRequest):
		return a.createProjectFromTextReply(turn, user, dialog)
	case matchWithRegexp(text, listProjiectsRequest):
		return a.listProjectsReply(turn, user, dialog)
	default:
		return a.defaultReply(turn, user, dialog)
	}
}

func (a *App) buildPayloadReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)
	payload := turn.Message.Payload.Value

	var header PayloadHeader
	err := json.Unmarshal([]byte(payload), &header)
	if err != nil {
		return reply, err
	}

	switch action := header.Action; action {
	case createProjectAction:
		return a.createProjectFromPayloadReply(turn, user, dialog)
	case showProjectsAction:
		return a.listProjectsReply(turn, user, dialog)
	case initCreateProjectAction:
		return a.initCreateProjectReply(turn, user, dialog)
	default:
		warn := fmt.Sprintf("got unknown action: %s", action)
		a.logger.Warn("failed to perform requested action", warn)
		reply.Message.Text = warn
		return reply, nil
	}
}

func (a *App) defaultReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)
	reply.Message.Text = defaultReplyText
	c := buildIntroCard()
	introCardJSON, err := c.StringIndent("", "  ")
	if err != nil {
		return reply, err
	}

	reply.Message.Attachment = introCardJSON
	return reply, nil
}

func (a *App) initCreateProjectReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)
	c := buildCreateProjectCard()
	createProjectCardJSON, err := c.StringIndent("", "  ")
	if err != nil {
		return reply, err
	}

	reply.Message.Attachment = createProjectCardJSON
	return reply, nil
}

func (a *App) createProjectFromTextReply(
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

	reply.Message.Text = buildCreateProjectSuccessMessage(project.Title, project.DueDate, project.ID)

	return reply, nil
}

func (a *App) createProjectFromPayloadReply(
	turn *mesga.Turn,
	user *domain.User,
	dialog *domain.Dialog,
) (*mesga.Turn, error) {
	reply := makeOutputTurn(turn)

	payload := turn.Message.Payload.Value
	var createProjectPayload CreateProjectPayload
	err := json.Unmarshal([]byte(payload), &createProjectPayload)
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}

	projectDueDate, err := time.Parse(defaultDateTimeLayout, createProjectPayload.DueDate)
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}

	project, err := a.CreateProject(user, &domain.Project{
		User:    user.Username,
		Title:   createProjectPayload.Title,
		DueDate: projectDueDate,
	})
	if err != nil {
		return nil, errors.New(buildCreateProjectFailedMessage(err))
	}

	reply.Message.Text = buildCreateProjectSuccessMessage(project.Title, project.DueDate, project.ID)

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
