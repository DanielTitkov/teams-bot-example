package mesga

import (
	"testing"
)

func sayOkFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	turn.Message.Text = "ok"
	return turn, nil
}

func defaultResponseFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	turn.Message.Text = "default response"
	return turn, nil
}

var setup = RouterSetup{
	GetStateFn:   getStateFn,
	StoreStateFn: storeStateFn,
	States: []StateSetup{
		{
			Title: "root",
			Actions: []ActionSetup{
				{
					TriggerPayloadAction: "sayOk",
					Function:             sayOkFn,
				},
				{
					TriggerText: "say ok",
					Function:    sayOkFn,
				},
				{
					TriggerTextRgxp: ".*say ok.*",
					Function:        sayOkFn,
				},
			},
			Default: ActionSetup{
				Function: defaultResponseFn,
			},
		},
	},
}

func TestRespondToPayload(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Payload: MessagePayload{
				Value: `{"action": "sayOk"}`,
			},
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := r.Respond(turn)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message.Text == "" {
		t.Fatal("got empty response")
	}

	if resp.Message.Text != "ok" {
		t.Errorf("expexted to get ok, got '%s'", resp.Message.Text)
	}
}

func TestRespondToUnknownPayloadAction(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Payload: MessagePayload{
				Value: `{"action": "wft"}`,
			},
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := r.Respond(turn); err == nil {
		t.Error("expected to get error, got nil")
	}
}

func TestRespondToText(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Text: "say ok",
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := r.Respond(turn)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message.Text == "" {
		t.Fatal("got empty response")
	}

	if resp.Message.Text != "ok" {
		t.Errorf("expexted to get ok, got '%s'", resp.Message.Text)
	}
}

func TestRespondToRegexp(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Text: "please say ok",
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := r.Respond(turn)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message.Text == "" {
		t.Fatal("got empty response")
	}

	if resp.Message.Text != "ok" {
		t.Errorf("expexted to get ok, got '%s'", resp.Message.Text)
	}
}

func TestRespondToRegexp2(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Text: "say ok, darling",
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := r.Respond(turn)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message.Text == "" {
		t.Fatal("got empty response")
	}

	if resp.Message.Text != "ok" {
		t.Errorf("expexted to get ok, got '%s'", resp.Message.Text)
	}
}

func TestDefaultResponse(t *testing.T) {
	turn := &Turn{
		Message: Message{
			Text: "wtf",
		},
	}

	r, err := NewRouter(setup)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := r.Respond(turn)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message.Text == "" {
		t.Fatal("got empty response")
	}

	if resp.Message.Text != "default response" {
		t.Errorf("expexted to get default response, got '%s'", resp.Message.Text)
	}
}
