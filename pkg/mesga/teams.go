package mesga

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

type (
	// Teams handles the HTTP requests from then connector service
	Teams struct {
		adapter          core.Adapter
		cfg              configs.Config
		logger           Logger
		onMessageHandler func(Turn) Turn
		onInvokeHandler  func(Turn) Turn
		onUpdateHandler  func(Turn) Turn
		sentChan         chan<- *Turn
	}
	// TeamsConfig holds data to create teams connector
	TeamsConfig struct {
		AppID            string          // AppID from Azure
		AppPassword      string          // AppPassword from Azure
		OnMessageHandler func(Turn) Turn // OnMessageHandler to process message activity
		OnInvokeHandler  func(Turn) Turn // OnInvokeHandler to process invoke activity
		OnUpdateHandler  func(Turn) Turn // OnUpdateHandler to process update activity
		sentChan         chan *Turn      // set automatically by manager
	}
)

func newTeams(
	logger Logger,
	cfg TeamsConfig,
) *Teams {
	setting := core.AdapterSetting{
		AppID:       cfg.AppID,
		AppPassword: cfg.AppPassword,
	}

	adapter, err := core.NewBotAdapter(setting)
	if err != nil {
		log.Fatal("error creating adapter: ", err)
	}

	return &Teams{
		adapter:          adapter,
		logger:           logger,
		onMessageHandler: cfg.OnMessageHandler,
		onInvokeHandler:  cfg.OnInvokeHandler,
		onUpdateHandler:  cfg.OnUpdateHandler,
	}
}

func (t *Teams) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var turn *Turn

	defer func() {
		turn.Message.Direction = OutputCode
		turn.Message.Proactive = false
		turn.System = TeamsCode
		t.turnToSentChan(turn)
	}()

	var handler = activity.HandlerFuncs{
		OnMessageFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onMessageHandler(t.activityToTurn(turnCtx))
			response.Related.Teams = turnCtx.Activity.ReplyToID
			turn = &response
			attachments, err := t.getAttachments(turn.Message.Attachment)
			if err != nil {
				turn.Err = err
				return schema.Activity{}, err
			}

			if response.DropRelated {
				var ref schema.ConversationReference
				err = json.Unmarshal([]byte(turn.Dialog.Teams), &ref)
				if err != nil {
					turn.Err = err
					return schema.Activity{}, err
				}
				err = t.adapter.DeleteActivity(context.TODO(), response.Related.Teams, ref)
				if err != nil {
					turn.Err = err
					return schema.Activity{}, err
				}
			}

			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text), activity.MsgOptionAttachments(attachments))
		},
		OnInvokeFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onInvokeHandler(t.activityToTurn(turnCtx))
			turn = &response
			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text))
		},
		OnConversationUpdateFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onUpdateHandler(t.activityToTurn(turnCtx))
			turn = &response
			return turnCtx.SendActivity(activity.MsgOptionText(response.Message.Text))
		},
	}

	act, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		t.logger.Error("failed to parse request", err)
		turn.Err = err
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.adapter.ProcessActivity(ctx, act, handler)
	if err != nil {
		t.logger.Error("failed to process request", err)
		turn.Err = err
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.logger.Info("request processed successfully", "")
}

func (t *Teams) getHandler() (func(w http.ResponseWriter, req *http.Request), error) {
	if t.onMessageHandler == nil {
		return nil, errors.New("message handler required")
	}

	if t.onInvokeHandler == nil {
		return nil, errors.New("invoke handler required")
	}

	if t.onUpdateHandler == nil {
		return nil, errors.New("update handler required")
	}

	return t.processMessage, nil
}

func (t *Teams) sendMessage(turn *Turn) *Turn {
	defer func() {
		turn.System = TeamsCode
		turn.Message.Direction = OutputCode
		turn.Message.Proactive = true
		t.turnToSentChan(turn)
	}()

	var ref schema.ConversationReference
	err := json.Unmarshal([]byte(turn.Dialog.Teams), &ref)
	if err != nil {
		t.logger.Error("failed to unmarshal conversation reference", err)
		turn.Err = err
		return turn
	}

	var proactiveHandler = activity.HandlerFuncs{
		OnMessageFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {

			attachments, err := t.getAttachments(turn.Message.Attachment)
			if err != nil {
				turn.Err = err
				return schema.Activity{}, err
			}

			return turnCtx.SendActivity(
				activity.MsgOptionText(turn.Message.Text),
				activity.MsgOptionAttachments(attachments),
			)
		},
	}

	err = t.adapter.ProactiveMessage(context.Background(), ref, proactiveHandler)
	if err != nil {
		t.logger.Error("Failed to send proactive message.", err)
		turn.Err = err
		return turn
	}
	t.logger.Info("Proactive message sent successfully.", fmt.Sprintf("%+v", turn.Message))
	return turn
}

func (t *Teams) activityToTurn(turnCtx *activity.TurnContext) Turn {
	conversationRef := activity.GetCoversationReference(turnCtx.Activity)
	jsonRef, err := json.Marshal(conversationRef)
	jsonValue, err := json.Marshal(turnCtx.Activity.Value)
	return Turn{
		Message: Message{
			Text: turnCtx.Activity.Text,
			Payload: MessagePayload{
				Type:  turnCtx.Activity.ValueType,
				Value: string(jsonValue),
			},
			Direction: InputCode,
		},
		Dialog: &Dialog{
			Teams: string(jsonRef),
		},
		User: &User{
			Teams: UserMessagerData{
				ID:       &conversationRef.User.ID,
				Username: &conversationRef.User.Name,
			},
		},
		System: TeamsCode,
		Err:    err,
	}
}

func (t *Teams) turnToSentChan(turn *Turn) {
	if t.sentChan == nil {
		t.logger.Warn("wasn't able to send turn to sent channel", "channel is not set")
		return
	}
	select {
	case t.sentChan <- turn:
	default:
		t.logger.Warn("wasn't able to send turn to sent channel", "default")
	}
}

func (t *Teams) getAttachments(rawAtt string) ([]schema.Attachment, error) {
	var attachments []schema.Attachment
	if rawAtt != "" {
		var attachment map[string]interface{}
		err := json.Unmarshal([]byte(rawAtt), &attachment)
		if err != nil {
			return attachments, err
		}
		attachments = []schema.Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				Content:     attachment,
			},
		}
	}
	return attachments, nil
}
