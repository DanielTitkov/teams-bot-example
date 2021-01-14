package main

import (
	"fmt"
	"log"

	"github.com/DanielTitkov/teams-bot-example/pkg/mesga/states"
)

func main() {
	log.Println("Running states")

	createProjectAction := states.Action{
		Title: "createProject",
	}
	addProjectTitleAction := states.Action{
		Title: "addProjectTitle",
	}

	rootState := states.State{
		Title:          "root",
		AllowedActions: []*states.Action{&createProjectAction},
	}

	createProjectState := states.State{
		Title:          "createProject",
		AllowedActions: []*states.Action{&addProjectTitleAction},
	}

	m := states.Manager{
		States: []*states.State{
			&rootState,
			&createProjectState,
		},
	}

	fmt.Println(m.Call(&addProjectTitleAction, ""))
}
