package mesga

type (
	// Turn container strores all recieved data
	Turn struct {
		Message Message
		Dialog  *Dialog
		User    *User
		System  string // message origin or messager to sent proactive
		Err     error
	}
	// User holds user data
	User struct {
		Teams    UserMessagerData
		Telegram UserMessagerData
		Slack    UserMessagerData
	}
	// UserMessagerData stores users data for a specific messager
	UserMessagerData struct {
		ID       *string
		Username *string
	}
	// Message is used for messagers
	Message struct {
		Text       string
		Attachment string
		Direction  string
		Proactive  bool
	}
	// Dialog holds dialog reference in different messagers
	Dialog struct {
		Teams    string // serialized reference
		Telegram string
		Slack    string
	}
)
