package slack

import (
	"github.com/sbstjn/hanu"
)

// Version is the current version of the Slack Bot.
const Version = "0.0.1"

// A Slack describes our slack integration bot service.
type Slack struct {
	bot *hanu.Bot
}

// New returns a new Slack bot.
func New() *Slack {
	return &Slack{
		bot: &hanu.Bot{},
	}
}

// Name returns the slack bot's name.
func (s *Slack) Name() string {
	return "slack"
}

// RegisterCommand registers a Command to the underline bot service.
//
// Should be called before `Listen`.
func (s *Slack) RegisterCommand(cmd Command) {
	text := cmd.Name
	if cmd.Args != "" {
		text += " " + cmd.Args
	}
	description := cmd.Description
	handler := func(c hanu.ConversationInterface) {
		cmd.Handler(s, c)
	}
	s.bot.Register(hanu.NewCommand(text, description, handler))
}

// Listen for message on underline socket.
func (s *Slack) Listen(token string) error {
	s.bot.Token = token
	if _, err := s.bot.Handshake(); err != nil {
		return err
	}

	for _, cmd := range ActiveCommands {
		s.RegisterCommand(cmd)
	}

	s.bot.Listen()
	return nil
}

// Shutdown does nothing here.
func (s *Slack) Shutdown() {

}
