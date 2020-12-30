package mesga

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	// TeamsCode is internal teams title
	TeamsCode = "teams"
	// SlackCode is internal slack title
	SlackCode = "slack"
	// TelegramCode is internal telegram title
	TelegramCode = "telegram"
)

type (
	// Manager is mesga general type to access all supported messagers
	Manager struct {
		teams         *Teams
		proactiveChan <-chan *Turn
		sentChan      chan<- *Turn
		logger        Logger
	}
	// Config has all setup data for mesga manager
	Config struct {
		Teams         TeamsConfig
		ProactiveChan chan *Turn // ProactiveChan accepts messages to be sent to user
		SentChan      chan *Turn // SentChan will be populated with messages that were sent
	}
)

// New creates messagers adapters and return new manager
func New(logger Logger, cfg Config) *Manager {
	teams := newTeams(logger, cfg.Teams)
	teams.sentChan = cfg.SentChan // pass channel to teams in order to get sent messages
	return &Manager{
		logger: logger,
		teams:  teams,
	}
}

// GetTeamsHandler returns http handle func to process teams requests
func (m *Manager) GetTeamsHandler() (func(w http.ResponseWriter, req *http.Request), error) {
	return m.teams.getHandler()
}

// RunProactiveListener starts listening proactive channel and sends messages that come to it
func (m *Manager) RunProactiveListener() error {
	if m.proactiveChan == nil {
		return errors.New("proactive channel is not set")
	}
	for {
		select {
		case turn := <-m.proactiveChan:
			switch turn.System {
			case TeamsCode:
				m.teams.sendMessage(turn)
			default:
				m.logger.Error("got unknown system code", fmt.Errorf("%s", turn.System))
			}
		default:
			m.logger.Warn("wasn't able to send proactive message", "default")
		}
	}
}
