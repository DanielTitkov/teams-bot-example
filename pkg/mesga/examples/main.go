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

	s := mesga.Setup(
		mesga.AddState(
			"root",
			mesga.AddAction("", createProjectFn),
			mesga.AddAction("", listProjectsFn),
		),
		mesga.AddState(
			"createProject",
			mesga.AddActionRgxp("", "", createProjectFn),
		),
	)

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
