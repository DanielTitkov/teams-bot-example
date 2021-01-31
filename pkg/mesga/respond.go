package mesga

import (
	"encoding/json"
	"errors"
)

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

	current, err := r.getState(turn)
	if err != nil {
		return turn, err
	}

	action, ok := current.payloadActionMapping[header.Action]
	if !ok {
		return turn, errors.New("got action that is not allowed in this state, got: " + payload)
	}

	reply, err := action.do(turn, current.Data)
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
