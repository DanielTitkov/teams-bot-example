package app

type (
	// PayloadHeader holds common fields to identify payload
	PayloadHeader struct {
		App    string
		Action string
	}
	// CreateProjectPayload holds data for new project
	CreateProjectPayload struct {
		PayloadHeader
		Title   string
		DueDate string
	}
)
