package mesga

import (
	"encoding/json"
	"errors"
)

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

func (r *Router) respondToPayload(turn *Turn) (*Turn, error) {
	// reply := makeOutputTurn(turn)
	payload := turn.Message.Payload.Value

	var header PayloadHeader
	err := json.Unmarshal([]byte(payload), &header)
	if err != nil {
		return turn, err
	}

	if header.Action == "" {
		return turn, errors.New("payload must have action, got: " + payload)
	}

	action, ok := r.Current.payloadActionMapping[header.Action]
	if !ok {
		return turn, errors.New("got action that is not allowed in this state, got: " + payload)
	}

	reply, err := action.do(turn, r.Current.Data)
	if err != nil {
		if action.OnFailTransition != nil {
			r.setState(action.OnFailTransition)
		}
		return reply, err
	}

	if action.OnSuccessTransition != nil {
		r.setState(action.OnSuccessTransition)
	}

	return reply, nil
}

func (r *Router) respondToText(turn *Turn) (*Turn, bool, error) {
	action, ok := r.Current.textActionMapping[turn.Message.Text]
	if !ok {
		return turn, false, nil // action not found, continue to regexp
	}

	reply, err := action.do(turn, r.Current.Data)
	if err != nil {
		if action.OnFailTransition != nil {
			r.setState(action.OnFailTransition)
		}
		return reply, true, err
	}

	if action.OnSuccessTransition != nil {
		r.setState(action.OnSuccessTransition)
	}

	return reply, true, nil
}

func (r *Router) respondToRegexp(turn *Turn) (*Turn, bool, error) {
	for regexp, action := range r.Current.textRgxpActionMapping {
		if regexp.Match([]byte(turn.Message.Text)) {
			reply, err := action.do(turn, r.Current.Data)
			if err != nil {
				if action.OnFailTransition != nil {
					r.setState(action.OnFailTransition)
				}
				return reply, true, err
			}

			if action.OnSuccessTransition != nil {
				r.setState(action.OnSuccessTransition)
			}

			return reply, true, nil
		}
	}

	return turn, false, nil // action not found
}

func (r *Router) respondDefault(turn *Turn) (*Turn, error) {
	reply, err := r.Current.defaultAction.do(turn, r.Current.Data)
	if err != nil {
		if r.Current.defaultAction.OnFailTransition != nil {
			r.setState(r.Current.defaultAction.OnFailTransition)
		}
		return reply, err
	}

	if r.Current.defaultAction.OnSuccessTransition != nil {
		r.setState(r.Current.defaultAction.OnSuccessTransition)
	}

	return reply, nil
}
