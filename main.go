// Package main is the commanad line tool for the ion (chat) bots.
//
// Usage: go run main.go -slack=the_secret_slack_bot_token
// Chat say: "@ion-bot version", replies to: "ion's current version is: 1.0.0".
package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/get-ion/bot/slack"
)

// BotService is the interface which should be completed by all bots integrations.
type BotService interface {
	// Name should return the name of the bot, will be used to take the token from the cli flags.
	Name() string
	// Listen should start listening based on the token.
	Listen(tok string) error
	// Shutdown should close the bot's connection.
	Shutdown()
}

func env(s string) string {
	osEnvKey := strings.ToUpper(s) + "_BOT_TOKEN" // SLACK_BOT_TOKEN
	return os.Getenv(osEnvKey)
}

func lookupToken(botName string) string {
	// search by cli flags
	flags := os.Args[1:]
	for _, s := range flags {
		if len(s) <= 2 {
			continue
		}
		// omit -
		// i.e: -slack=...
		if s[0] == '-' {
			s = s[1:]
		}
		// split name from value
		// i.e: slack=val-val-val to name = [slack], value = [val-val-val]
		if strings.Contains(s, "=") {
			splitted := strings.Split(s, "=")
			if len(splitted) >= 2 {
				// if name matches
				if name := splitted[0]; name == botName {
					// return the token value
					value := splitted[1]
					return value
				}
			}

		}
	}
	// search by system env
	return env(botName)
}

var bots = []BotService{
	slack.New(),
}

func main() {
	for _, b := range bots {
		name := b.Name()
		token := lookupToken(name)
		if token != "" {
			log.Printf("%s bot is listening\n", name)
			go b.Listen(token)
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, os.Kill, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-ch:
		log.Println("exiting...")
		for _, b := range bots {
			b.Shutdown()
		}

		os.Exit(0)
	}
}
