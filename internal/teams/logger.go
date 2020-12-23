package teams

// Logger is a general logger interface for teams
type Logger interface {
	Info(string, string)
	Warn(string, string)
	Error(string, error)
}
