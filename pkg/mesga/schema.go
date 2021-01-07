package mesga

type (
	// Turn container strores all recieved data
	Turn struct {
		Message     Message
		Dialog      *Dialog
		User        *User
		Related     Related
		System      string // message origin or messager to sent proactive
		Err         error
		DropRelated bool // tells to delete related message (for constructors etc)
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
		Payload    MessagePayload
	}
	// MessagePayload stores teams activity value staff
	MessagePayload struct {
		Type  string
		Value string // serialized payload // TODO: maybe have it as a map?
	}
	// Dialog holds dialog reference in different messagers
	Dialog struct {
		Teams    string // serialized reference
		Telegram string
		Slack    string
	}
	// Related holds reference to related message
	Related struct {
		Teams    string // serialized reference
		Telegram string
		Slack    string
	}
)
