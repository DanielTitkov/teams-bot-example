package teams

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/DanielTitkov/teams-bot-example/internal/configs"
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/logger"
	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

const (
	systemCode = "teams"
	ouputCode  = "output"
	inputCode  = "input"
)

// Teams handles the HTTP requests from then connector service
type Teams struct {
	adapter          core.Adapter
	cfg              configs.Config
	logger           Logger
	onMessageHandler func(domain.Turn) domain.Turn
	onInvokeHandler  func(domain.Turn) domain.Turn
	onUpdateHandler  func(domain.Turn) domain.Turn
	proactiveChan    <-chan *domain.Turn
	sentChan         chan<- *domain.Turn
}

func NewTeams(
	cfg configs.Config,
	logger *logger.Logger,
) *Teams {
	setting := core.AdapterSetting{
		AppID:       cfg.Teams.AppID,
		AppPassword: cfg.Teams.AppPassword,
	}

	adapter, err := core.NewBotAdapter(setting)
	if err != nil {
		log.Fatal("error creating adapter: ", err)
	}

	return &Teams{
		adapter: adapter,
		cfg:     cfg,
		logger:  logger,
	}
}

func (t *Teams) SetOnMessageHandler(handler func(domain.Turn) domain.Turn) {
	t.onMessageHandler = handler
}

func (t *Teams) SetOnInvokeHandler(handler func(domain.Turn) domain.Turn) {
	t.onInvokeHandler = handler
}

func (t *Teams) SetOnUpdateHandler(handler func(domain.Turn) domain.Turn) {
	t.onUpdateHandler = handler
}

func (t *Teams) SetProactiveChannel(ch chan *domain.Turn) {
	t.proactiveChan = ch
}

func (t *Teams) SetSentChannel(ch chan *domain.Turn) {
	t.sentChan = ch
}

func (t *Teams) processMessage1(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var turn *domain.Turn

	defer func() {
		turn.Message.Direction = ouputCode
		turn.Message.System = systemCode
		turn.Message.Proactive = false
		t.turnToSentChan(turn)
	}()

	var handler = activity.HandlerFuncs{
		OnMessageFunc: func(turnCtx *activity.TurnContext) (schema.Activity, error) {
			response := t.onMessageHandler(t.activityToTurn(turnCtx))
			turn = &response
			attachments, err := t.getAttachments(turn.Message.Attachment)
			if err != nil {
				turn.Err = err
				return schema.Activity{}, err
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

func (t *Teams) Listen() error {
	if t.onMessageHandler == nil {
		return errors.New("message handler required")
	}

	if t.onInvokeHandler == nil {
		return errors.New("invoke handler required")
	}

	if t.onUpdateHandler == nil {
		return errors.New("update handler required")
	}

	http.HandleFunc("/api/messages", t.processMessage)
	port := fmt.Sprintf(":%d", t.cfg.Teams.Port)
	t.logger.Info("Listening for teams messages", "port "+port)
	http.ListenAndServe(port, nil)
	return nil
}

func (t *Teams) RunProactiveManager() {
	for {
		select {
		case message := <-t.proactiveChan:
			t.sendMessage(message)
		}
	}
}

func (t *Teams) sendMessage(turn *domain.Turn) *domain.Turn {
	defer func() {
		turn.Message.Direction = ouputCode
		turn.Message.System = systemCode
		turn.Message.Proactive = true
		t.turnToSentChan(turn)
	}()

	var ref schema.ConversationReference
	err := json.Unmarshal([]byte(turn.Dialog.Meta.Teams), &ref)
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

func (t *Teams) activityToTurn(turnCtx *activity.TurnContext) domain.Turn {
	conversationRef := activity.GetCoversationReference(turnCtx.Activity)
	jsonRef, err := json.Marshal(conversationRef)
	return domain.Turn{
		Message: domain.Message{
			Text: turnCtx.Activity.Text,
		},
		Dialog: domain.TurnDialog{
			Meta: domain.DialogMeta{
				Teams: string(jsonRef),
			},
		},
		User: domain.TurnUser{
			Meta: domain.UserMeta{
				Teams: domain.UserMessagerData{
					ID:       &conversationRef.User.ID,
					Username: &conversationRef.User.Name,
				},
			},
		},
		Err: err,
	}
}

func (t *Teams) turnToSentChan(turn *domain.Turn) {
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

// file

func putRequest(u string, data []byte) error {
	client := &http.Client{}
	dec, err := url.QueryUnescape(u)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, dec, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	size := fmt.Sprintf("%d", len(data))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-Length", size)
	req.Header.Set("Content-Range", fmt.Sprintf("bytes 0-%d/%d", len(data)-1, len(data)))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return fmt.Errorf("failed to upload file with status %d", resp.StatusCode)
	}
	return nil
}

func (t *Teams) cleanupConsents(ID string, ref schema.ConversationReference) {
	fmt.Printf("Deleting activity %s\n", ID)
	if err := t.adapter.DeleteActivity(context.TODO(), ID, ref); err != nil {
		log.Printf("Failed to delete activity. %s", err.Error())
	}
}

func (t *Teams) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	act, err := t.adapter.ParseRequest(ctx, req)
	if err != nil {
		fmt.Println("Failed to parse request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aj, _ := json.MarshalIndent(act, "", "  ")
	fmt.Printf("Incoming Activity:: \n%s\n", aj)

	customHandler := activity.HandlerFuncs{
		// handle message events
		OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			fi, err := os.Stat("data.txt")
			if err != nil {
				return schema.Activity{}, fmt.Errorf("failed to read file: %s", err.Error())
			}

			// send file upload request
			attachments := []schema.Attachment{
				{
					ContentType: "application/vnd.microsoft.teams.card.file.consent",
					Name:        "data.txt",
					Content: map[string]interface{}{
						"description": "Sample data",
						"sizeInBytes": fi.Size(),
					},
				},
			}
			return turn.SendActivity(activity.MsgOptionText("Echo: "+turn.Activity.Text), activity.MsgOptionAttachments(attachments))
		},
		// handle invoke events
		// https://developer.microsoft.com/en-us/microsoft-teams/blogs/working-with-files-in-your-microsoft-teams-bot/
		OnInvokeFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
			t.cleanupConsents(turn.Activity.ReplyToID, activity.GetCoversationReference(turn.Activity))
			data, err := ioutil.ReadFile("data.txt")
			if err != nil {
				return schema.Activity{}, fmt.Errorf("failed to read file: %s", err.Error())
			}
			if turn.Activity.Value["type"] != "fileUpload" {
				return schema.Activity{}, nil
			}
			if turn.Activity.Value["action"] != "accept" {
				return schema.Activity{}, nil
			}

			// parse upload info from invoke accept response
			uploadInfo := schema.UploadInfo{}
			infoJSON, err := json.Marshal(turn.Activity.Value["uploadInfo"])
			if err != nil {
				return schema.Activity{}, err
			}
			err = json.Unmarshal(infoJSON, &uploadInfo)
			if err != nil {
				return schema.Activity{}, err
			}

			// upload file
			err = putRequest(uploadInfo.UploadURL, data)
			if err != nil {
				return schema.Activity{}, fmt.Errorf("failed to upload file: %s", err.Error())
			}

			// notify user about uploaded file
			fileAttach := []schema.Attachment{
				{
					ContentType: "application/vnd.microsoft.teams.card.file.info",
					ContentURL:  uploadInfo.ContentURL,
					Name:        uploadInfo.Name,
					Content: map[string]interface{}{
						"uniqueId": uploadInfo.UniqueID,
						"fileType": uploadInfo.FileType,
					},
				},
			}

			return turn.SendActivity(activity.MsgOptionAttachments(fileAttach))
		},
	}

	err = t.adapter.ProcessActivity(ctx, act, customHandler)
	if err != nil {
		fmt.Println("Failed to process request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Request processed successfully.")
}
