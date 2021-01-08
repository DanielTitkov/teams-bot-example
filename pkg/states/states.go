package states

import (
	"fmt"
	"sync"
)

// Manager is a state engine that manages dialog states
type Manager struct {
	States  []*State
	Root    *State
	Current *State

	mx                sync.Mutex
	transitionMapping map[string]string                     // or maybe state to state
	useMapping        bool                                  // otherwise use states' transitions
	getStateFn        func(jsonArgs string) (*State, error) // for external state storage
	storeStateFn      func(*State) error                    // for external state storage
}

// Call tries to perform some action
func (m *Manager) Call(action *Action, jsonArgs string) (jsonResult string, err error) {
	m.mx.Lock()
	defer m.mx.Unlock()

	current, err := m.getState(jsonArgs)
	if err != nil {
		return jsonResult, err
	}

	if !isAllowed(action, current.AllowedActions) {
		return jsonResult, fmt.Errorf("action %s is not allowed in state %s", action.Title, m.Current.Title)
	}

	res, err := action.do(jsonArgs, current.Data)
	if err != nil {
		if action.OnFailTransition != nil {
			m.setState(action.OnFailTransition)
		}
		return jsonResult, err
	}

	if action.OnSuccessTransition != nil {
		m.setState(action.OnSuccessTransition)
	}

	return res, nil
}

func (m *Manager) setState(state *State) {
	m.mx.Lock()
	defer m.mx.Unlock()
	if m.storeStateFn != nil {
		m.storeStateFn(state) // TODO handle error
	}
	m.Current = state
}

func (m *Manager) getState(jsonArgs string) (*State, error) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.getStateFn != nil {
		current, err := m.getStateFn(jsonArgs)
		if err != nil {
			return nil, err
		}
		return current, nil
	}
	current := m.Current
	return current, nil
}

// State is dialog state
type State struct {
	Title            string
	Data             map[string]interface{}
	AllowedActions   []*Action
	OnEnterCallbacks []*Action
	OnExitCallbacks  []*Action
	// IsRoot           bool
}

// Action is something to do within state
type Action struct {
	Title               string
	OnSuccessTransition *State
	OnFailTransition    *State

	fn func(jsonArgs string, data map[string]interface{}) (jsonResult string, err error)
}

func (a *Action) do(jsonArgs string, data map[string]interface{}) (jsonResult string, err error) {
	res, err := a.fn(jsonArgs, data)
	if err != nil {
		return jsonResult, err
	}
	return res, nil
}

func isAllowed(action *Action, allowed []*Action) bool {
	for _, a := range allowed {
		if a.Title == action.Title {
			return true
		}
	}
	return false
}
