package domain

import "time"

const ()

type (
	// User holds user data
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service
		Username     string
		DisplayName  string // human readable name
		Password     string
		PasswordHash string
		Email        string   // TODO: add ent validation
		Meta         UserMeta // stores messagers ids
	}
	// Message is used for messagers
	Message struct { // TODO: maybe create container objects instead
		ID         int
		Text       string
		Attachment string
		System     string
		Direction  string
		Proactive  bool
	}
	// Turn container strores all recieved data
	Turn struct {
		Message Message
		Dialog  TurnDialog // used for callbacks, not stored
		User    TurnUser   // TODO: think about it
		Err     error
	}
	// TurnUser hodls dialog data for message processor
	TurnUser struct {
		User *User
		Meta UserMeta
	}
	// TurnDialog hodls dialog data for message processor
	TurnDialog struct {
		Dialog *Dialog
		Meta   DialogMeta
	}
	// UserMeta stores users data for messagers
	UserMeta struct {
		Teams    UserMessagerData
		Telegram UserMessagerData
		Slack    UserMessagerData
	}
	// UserMessagerData stores users data for a specific messager
	UserMessagerData struct {
		ID       *string
		Username *string
	}
	// Dialog holds dialog reference in different messagers
	Dialog struct {
		ID   int
		User string // username
		Meta DialogMeta
	}
	// DialogMeta stores serialize dialog reference for messagers
	DialogMeta struct {
		Teams    string
		Telegram string // for example
		Slack    string // for example
	}
	// Business logic
	// not linked to messenges staff
	// Project
	Project struct {
		ID      int
		User    string
		Title   string
		DueDate time.Time
	}
)
