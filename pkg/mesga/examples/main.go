package main

import (
	"fmt"
	"log"

	"github.com/DanielTitkov/teams-bot-example/pkg/mesga"
)

func createProjectFn(turn *mesga.Turn, data map[string]interface{}) (reply *mesga.Turn, err error) {
	fmt.Println("running create project")
	return turn, nil
}

func addProjectTitleFn(turn *mesga.Turn, data map[string]interface{}) (reply *mesga.Turn, err error) {
	fmt.Println("running add project title")
	return turn, nil
}

func listProjectsFn(turn *mesga.Turn, data map[string]interface{}) (reply *mesga.Turn, err error) {
	fmt.Println("running list projects")
	return turn, nil
}

func main() {
	log.Println("Running mesga")

	setup := mesga.RouterSetup{
		States: []mesga.StateSetup{
			{
				Title: "root",
				Actions: []mesga.ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    createProjectFn,
					},
					{
						TriggerText:         "Мои проекты",
						OnSuccessTransition: "createProject",
						Function:            listProjectsFn,
					},
				},
				Default: mesga.ActionSetup{},
				Data:    map[string]interface{}{},
			},
			{
				Title: "createProject",
				Actions: []mesga.ActionSetup{
					{
						TriggerTextRgxp:     ".*",
						OnSuccessTransition: "root",
						OnFailTransition:    "root",
						Function:            addProjectTitleFn,
					},
				},
				Default: mesga.ActionSetup{},
				Data:    map[string]interface{}{},
			},
		},
	}

	fmt.Printf("\nSETUP\n%+v\n", setup)
	r, err := mesga.NewRouter(setup)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nROUTER\n%+v\n", r)

	// TODO: test responders
}
