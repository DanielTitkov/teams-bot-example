package mesga

import (
	"errors"
	"regexp"
	"sync"
)

// Router is a state engine that manages dialog states
type Router struct {
	Root    *State
	Current *State

	stateMapping      map[string]*State
	mx                sync.Mutex
	transitionMapping map[string]string           // or maybe state to state
	useMapping        bool                        // otherwise use states' transitions
	getStateFn        func(*Turn) (*State, error) // for external state storage
	storeStateFn      func(*State) error          // for external state storage
}

// Respond generates new turn
func (r *Router) Respond(turn *Turn) (*Turn, error) {
	current, err := r.getState(turn)
	if err != nil {
		return turn, err
	}
	r.mx.Lock()
	r.Current = current
	r.mx.Unlock() // FIXME

	if turn.Message.Payload.Value != "" {
		response, err := r.respondToPayload(turn)
		return response, err // TODO: maybe process error in turn
	}

	response, ok, err := r.respondToText(turn)
	if ok {
		if err != nil {
			return response, err
		}
		return response, nil
	}

	response, ok, err = r.respondToRegexp(turn)
	if ok {
		if err != nil {
			return response, err
		}
		return response, nil
	}

	if current.defaultAction != nil {
		defaultResponse, err := r.respondDefault(turn)
		if err != nil {
			return defaultResponse, err
		}
	}

	return turn, nil
}

func (m *Router) setState(state *State) {
	m.mx.Lock()
	defer m.mx.Unlock()
	if m.storeStateFn != nil {
		m.storeStateFn(state) // TODO handle error
	}
	m.Current = state
}

func (m *Router) getState(turn *Turn) (*State, error) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.getStateFn != nil {
		current, err := m.getStateFn(turn)
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
	Title                 string
	Data                  map[string]interface{}
	AllowedActions        []*Action
	OnEnterCallbacks      []*Action
	OnExitCallbacks       []*Action
	textActionMapping     map[string]*Action
	textRgxpActionMapping map[*regexp.Regexp]*Action
	payloadActionMapping  map[string]*Action
	defaultAction         *Action
}

// Action is something to do within state
type Action struct {
	Title               string
	OnSuccessTransition *State
	OnFailTransition    *State

	fn func(turn *Turn, data map[string]interface{}) (reply *Turn, err error)
}

func (a *Action) do(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	if a.fn == nil {
		return reply, errors.New("action fn is not set")
	}
	res, err := a.fn(turn, data)
	if err != nil {
		return reply, err
	}
	return res, nil
}
