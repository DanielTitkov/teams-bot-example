package mesga

import (
	"errors"
	"fmt"
)

func NewRouter(setup RouterSetup) (*Router, error) {
	stateMapping := make(map[string]*State)

	var states []*State
	for _, stateSetup := range setup.States {
		state := &State{
			Title: stateSetup.Title,
			Data:  stateSetup.Data,
			// textActionMapping:     textActionMapping,
			// textRgxpActionMapping: textRgxpActionMapping,
			// payloadActionMapping:  payloadActionMapping,
		}

		if _, ok := stateMapping[stateSetup.Title]; ok {
			return nil, errors.New("state title must be unique, but title is repeated: " + stateSetup.Title)
		}
		stateMapping[stateSetup.Title] = state
	}

	// after states are created populate them with actions
	// two loops are needed for validation
	// for action can be validated on when all states are known
	for _, stateSetup := range setup.States {
		// state := stateMapping[stateSetup.Title]

		// textActionMapping := make(map[string]*Action)
		// textRgxpActionMapping := make(map[string]*Action)
		// payloadActionMapping := make(map[string]*Action)

		for _, actionSetup := range stateSetup.Actions {
			if actionSetup.OnSuccessTransition != "" {
				if _, ok := stateMapping[actionSetup.OnSuccessTransition]; !ok {
					return nil, errors.New("action declares transition to an unknown state: " + actionSetup.OnSuccessTransition)
				}
			}
			if actionSetup.OnFailTransition != "" {
				if _, ok := stateMapping[actionSetup.OnFailTransition]; !ok {
					return nil, errors.New("action declares transition to an unknown state: " + actionSetup.OnFailTransition)
				}
			}

			// validate action has only one trigger
			if !xor(
				actionSetup.TriggerText != "",
				actionSetup.TriggerTextRgxp != "",
				actionSetup.TriggerPayloadAction != "",
			) {
				return nil, errors.New("action must have only one trigger type")
			}

			// action := &Action{}
		}

	}

	router := &Router{
		States: states,
	}

	return router, nil
}

type RouterSetup struct {
	States []StateSetup
}

type StateSetup struct {
	Title            string
	Actions          []ActionSetup
	Default          ActionSetup
	Data             map[string]interface{} // default data
	OnEnterCallbacks []func() error
	OnExitCallbacks  []func() error
}

type ActionSetup struct {
	TriggerText          string // or
	TriggerTextRgxp      string // or
	TriggerPayloadAction string
	OnSuccessTransition  string
	OnFailTransition     string
	Function             func(*Turn, map[string]interface{}) (*Turn, error)
}

func xor(bools ...bool) bool {
	var trues []struct{}
	for _, b := range bools {
		if b {
			trues = append(trues, struct{}{})
		}
	}
	fmt.Println(trues)
	if len(trues) == 1 {
		return true
	}
	return false
}
