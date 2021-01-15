package main

import (
	"fmt"
	"log"

	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func createProjectFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	fmt.Println("running create project")
	return turn, nil
}

func addProjectTitleFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	fmt.Println("running add project title")
	return turn, nil
}

func listProjectsFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	fmt.Println("running list projects")
	return turn, nil
}

func main() {
	log.Println("Running mesga")

	r := mesga.RouterSetup{
		States: []mesga.StateSetup{
			{
				Title: "root",
				Actions: []mesga.ActionSetup{
					{
						TriggerText:          "Создать проект",
						TriggerTextRgxp:      "",
						TriggerPayloadAction: "",
						OnSuccessTransition:  "",
						OnFailTransition:     "",
						Function:             createProjectFn,
					},
					{
						TriggerText:          "Мои проекты",
						TriggerTextRgxp:      "",
						TriggerPayloadAction: "",
						OnSuccessTransition:  "createProject",
						OnFailTransition:     "",
						Function:             listProjectsFn,
					},
				},
				Default: mesga.ActionSetup{},
				Data:    map[string]interface{}{},
			},
			{
				Title: "createProject",
				Actions: []mesga.ActionSetup{
					{
						TriggerText:          "",
						TriggerTextRgxp:      ".*",
						TriggerPayloadAction: "",
						OnSuccessTransition:  "root",
						OnFailTransition:     "root",
						Function:             addProjectTitleFn,
					},
				},
				Default: mesga.ActionSetup{},
				Data:    map[string]interface{}{},
			},
		},
	}

	createProjectAction := mesga.Action{
		Title: "createProject",
	}
	addProjectTitleAction := mesga.Action{
		Title: "addProjectTitle",
	}

	rootState := mesga.State{
		Title:          "root",
		AllowedActions: []*mesga.Action{&createProjectAction},
	}

	createProjectState := mesga.State{
		Title:          "createProject",
		AllowedActions: []*mesga.Action{&addProjectTitleAction},
	}

	m := mesga.States{
		States: []*mesga.State{
			&rootState,
			&createProjectState,
		},
		Root:    &rootState,
		Current: &rootState,
	}

	fmt.Println(m.Call(&addProjectTitleAction, &mesga.Turn{}))
	fmt.Println(m.Call(&createProjectAction, &mesga.Turn{}))
}
