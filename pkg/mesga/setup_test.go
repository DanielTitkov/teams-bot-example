package mesga

import (
	"fmt"
	"testing"
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

func TestNonUniqueStates(t *testing.T) {
	setup := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerTextRgxp: ".*",
						Function:        addProjectTitleFn,
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup); err == nil {
		t.Error("expected to get an error, but got nil")
	}
}

func TestIncorrectTransitions(t *testing.T) {
	setup := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText:         "Создать проект",
						Function:            createProjectFn,
						OnSuccessTransition: "state2",
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup); err == nil {
		t.Error("expected to get an error, but got nil")
	}

	setup2 := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText:      "Создать проект",
						Function:         createProjectFn,
						OnFailTransition: "state2",
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup2); err == nil {
		t.Error("expected to get an error, but got nil")
	}
}

func TestMulptipleTriggers(t *testing.T) {
	setup := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText:     "Создать проект",
						TriggerTextRgxp: ".*",
						Function:        createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup); err == nil {
		t.Error("expected to get an error, but got nil")
	}

	setup2 := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText:          "Создать проект",
						TriggerPayloadAction: "action",
						Function:             createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup2); err == nil {
		t.Error("expected to get an error, but got nil")
	}

	setup3 := RouterSetup{
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerTextRgxp:      "Создать проект",
						TriggerPayloadAction: "action",
						Function:             createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup3); err == nil {
		t.Error("expected to get an error, but got nil")
	}
}
