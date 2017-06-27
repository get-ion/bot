package slack

import (
	"github.com/sbstjn/hanu"
)

// Handler is the handler signature for a Command.
type Handler func(bot *Slack, c hanu.ConversationInterface)

// A Command describes a bot's command.
//
// It's "Name" and "Handler" are required.
type Command struct {
	Name        string // hello
	Description string
	Args        string // <word>
	Handler     Handler
}

// ActiveCommands the commands that are enabled and ready to being used.
var ActiveCommands = []Command{
	GopherNewbieResourcesCommand,
	GopherRecommendedChannelsCommand,
	GopherDatabaseTutorialCommand,
	GopherPackageLayoutCommand,
	GopherLibraryForCommand,
	VersionCommand,
	VersionBotCommand,
}
