package mesga

import (
	"errors"
	"regexp"
)

func NewRouter(setup RouterSetup) (*Router, error) {
	if setup.StoreStateFn == nil {
		return nil, errors.New("router must have store state function") // TODO: maybe add default funcs
	}

	if setup.GetStateFn == nil {
		return nil, errors.New("router must have get state function")
	}

	stateMapping := make(map[string]*State)

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

	root, ok := stateMapping["root"]
	if !ok {
		return nil, errors.New("root state is required")
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

		// set state default action
		if stateSetup.Default.Function != nil {
			if err := stateSetup.Default.validate(stateMapping, false); err != nil {
				return nil, err
			}
			state.defaultAction = &Action{
				OnSuccessTransition: stateMapping[stateSetup.Default.OnSuccessTransition],
				OnFailTransition:    stateMapping[stateSetup.Default.OnFailTransition],
				fn:                  stateSetup.Default.Function,
			}
		}

		// set state actions
		for _, actionSetup := range stateSetup.Actions {
			if err := actionSetup.validate(stateMapping, true); err != nil {
				return nil, err
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
		// add state callbacks?
		// add systems
		// add proactive stack?
	}

	router := &Router{
		Root:         root,
		Current:      root,
		stateMapping: stateMapping,
		getStateFn:   setup.GetStateFn,
		storeStateFn: setup.StoreStateFn,
	}

	router.setState(root)

	return router, nil
}

// RouterSetup holds options to build router
type RouterSetup struct {
	States       []StateSetup
	GetStateFn   func(*Turn) (*State, error)
	StoreStateFn func(*State) error
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

func (s *ActionSetup) validate(stateMapping map[string]*State, validateTriggers bool) error {
	if s.Function == nil {
		return errors.New("action must have function")
	}

	if s.OnSuccessTransition != "" {
		if _, ok := stateMapping[s.OnSuccessTransition]; !ok {
			return errors.New("action declares transition to an unknown state: " + s.OnSuccessTransition)
		}
	}
	if s.OnFailTransition != "" {
		if _, ok := stateMapping[s.OnFailTransition]; !ok {
			return errors.New("action declares transition to an unknown state: " + s.OnFailTransition)
		}
	}

	// validate action has only one trigger
	if validateTriggers && !xor(
		s.TriggerText != "",
		s.TriggerTextRgxp != "",
		s.TriggerPayloadAction != "",
	) {
		return errors.New("action must have exactly one trigger")
	}
	return nil
}
