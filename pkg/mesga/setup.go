package mesga

import (
	"errors"
	"regexp"
)

func NewRouter(setup RouterSetup) (*Router, error) {
	stateMapping := make(map[string]*State)

	var states []*State
	for _, stateSetup := range setup.States {
		state := &State{
			Title: stateSetup.Title,
			Data:  stateSetup.Data,
		}

		if _, ok := stateMapping[stateSetup.Title]; ok {
			return nil, errors.New("state title must be unique, but title is repeated: " + stateSetup.Title)
		}
		stateMapping[stateSetup.Title] = state
	}

	// after states are created populate them with actions
	// two loops are needed for validation
	// for actions can be validated only when all states are known
	for _, stateSetup := range setup.States {
		state := stateMapping[stateSetup.Title]

		textActionMapping := make(map[string]*Action)
		payloadActionMapping := make(map[string]*Action)
		textRgxpActionMapping := make(map[*regexp.Regexp]*Action)
		regexpMapping := make(map[string]struct{}) // for regexp validaion

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
				return nil, errors.New("action must have exactly one trigger")
			}

			action := &Action{
				OnSuccessTransition: stateMapping[actionSetup.OnSuccessTransition],
				OnFailTransition:    stateMapping[actionSetup.OnFailTransition],
				fn:                  actionSetup.Function,
			}

			if actionSetup.TriggerText != "" {
				if _, ok := textActionMapping[actionSetup.TriggerText]; ok {
					return nil, errors.New("action triggers must be unique within a state")
				}
				textActionMapping[actionSetup.TriggerText] = action
			}

			if actionSetup.TriggerPayloadAction != "" {
				if _, ok := payloadActionMapping[actionSetup.TriggerPayloadAction]; ok {
					return nil, errors.New("action triggers must be unique within a state")
				}
				payloadActionMapping[actionSetup.TriggerPayloadAction] = action
			}

			if actionSetup.TriggerTextRgxp != "" {
				if _, ok := regexpMapping[actionSetup.TriggerTextRgxp]; ok {
					return nil, errors.New("action triggers must be unique within a state")
				}
				regexpMapping[actionSetup.TriggerTextRgxp] = struct{}{}
				rgxp, err := regexp.Compile(actionSetup.TriggerTextRgxp)
				if err != nil {
					return nil, err
				}
				textRgxpActionMapping[rgxp] = action
			}

			state.payloadActionMapping = payloadActionMapping
			state.textActionMapping = textActionMapping
			state.textRgxpActionMapping = textRgxpActionMapping
		}
		// TODO:
		// add state functions
		// add root and current
		// add state callbacks
		// validate action has function
		// add systems
		// add proactive stack?
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
	if len(trues) == 1 {
		return true
	}
	return false
}
