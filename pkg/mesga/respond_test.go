package mesga

import (
	"testing"
)

func sayOkFn(turn *Turn, data map[string]interface{}) (reply *Turn, err error) {
	turn.Message.Text = "ok"
	return turn, nil
}

func TestRespondToPayload(t *testing.T) {
	setup := RouterSetup{
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
				},
				Default: ActionSetup{},
			},
		},
	}

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
		t.Errorf("expexted to got ok, got '%s'", resp.Message.Text)
	}

}

func TestRespondToUnknownPayloadAction(t *testing.T) {
	setup := RouterSetup{
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
				},
				Default: ActionSetup{},
			},
		},
	}

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
		t.Error("expected to have error, got nil")
	}
}
