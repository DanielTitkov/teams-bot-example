package domain

const ()

type (
	// User holds user data
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
	}
	// Message is used for messagers
	Message struct {
		Text       string
		DialogData string
		Attachment string
	}
	// Dialog holds dialog reference in different messagers
	Dialog struct {
		ID   int
		User string // username
		Meta DialogMeta
	}
	// DialogMeta for messagers
	DialogMeta struct {
		Teams    string
		Telegram string // for example
		Slack    string // for example
	}
)
