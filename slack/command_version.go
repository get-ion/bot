package slack

import (
	"github.com/get-ion/ion"

	"github.com/sbstjn/hanu"
)

// VersionCommand "version" shows the ion's current version.
var VersionCommand = Command{
	Name: "version",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply("ion's current version is: " + ion.Version)
	},
}

// BotVersionCommand "bot-version" shows the bot's current version.
var BotVersionCommand = Command{
	Name: "bot-version",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply("bot's current version is: " + Version)
	},
}
