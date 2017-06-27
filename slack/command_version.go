package slack

import (
	"github.com/get-ion/ion"

	"github.com/sbstjn/hanu"
)

// VersionCommand "version" shows the ion's current version.
var VersionCommand = Command{
	Name:        "version",
	Description: "outputs the ion's current version",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply("ion's current version is: " + ion.Version)
	},
}

// VersionBotCommand "version bot" shows the bot's current version.
var VersionBotCommand = Command{
	Name:        "version bot",
	Description: "outputs the bot's current version",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply("bot's current version is: " + Version)
	},
}
