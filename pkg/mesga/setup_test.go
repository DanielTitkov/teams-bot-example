package mesga

import (
	"fmt"
	"testing"
)

var storedState *State

func storeStateFn(state *State) error {
	storedState = state
	return nil
}

func getStateFn(turn *Turn) (*State, error) {
	return storedState, nil
}

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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
			{
				Title: "root",
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

func TestActionNoFunction(t *testing.T) {
	setup := RouterSetup{
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
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

func TestNoRootState(t *testing.T) {
	setup := RouterSetup{
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "state1",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    addProjectTitleFn,
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

func TestNoStateFunctions(t *testing.T) {
	setup := RouterSetup{
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    addProjectTitleFn,
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
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

func TestNoTriggers(t *testing.T) {
	setup := RouterSetup{
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						Function: createProjectFn,
					},
				},
				Default: ActionSetup{},
			},
		},
	}

	if _, err := NewRouter(setup); err == nil {
		t.Error("expected to get an error, but got nil") // TODO: check error type
	}
}

func TestRepeatedTriggers(t *testing.T) {
	setup := RouterSetup{
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerText: "Создать проект",
						Function:    createProjectFn,
					},
					{
						TriggerText: "Создать проект",
						Function:    createProjectFn,
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerPayloadAction: "action",
						Function:             createProjectFn,
					},
					{
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
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerTextRgxp: ".*",
						Function:        createProjectFn,
					},
					{
						TriggerTextRgxp: ".*",
						Function:        createProjectFn,
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

func TestBadRgxpTrigger(t *testing.T) {
	setup := RouterSetup{
		GetStateFn:   getStateFn,
		StoreStateFn: storeStateFn,
		States: []StateSetup{
			{
				Title: "root",
				Actions: []ActionSetup{
					{
						TriggerTextRgxp: "***",
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
}
