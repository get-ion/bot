package slack

import (
	"fmt"

	"github.com/sbstjn/hanu"
)

/*
Contains some of the "gopher" bot's commands taken from the https://gophers.slack.com/messages/@gopher/.
*/

const gopherNewbieResourcesReply = `Here are some resources you should check out if you are learning / new to Go:
First you should take the language tour: http://tour.golang.org/

Then, you should visit:
- https://golang.org/doc/code.html to learn how to organize your Go workspace
- https://golang.org/doc/effective_go.html be more effective at writing Go
- https://golang.org/ref/spec learn more about the language itself
- https://golang.org/doc/#articles a lot more reading material

There are some awesome websites as well:
- https://blog.gopheracademy.com great resources for Gophers in general
- http://gotime.fm awesome weekly podcast of Go awesomeness
- https://gobyexample.com examples of how to do things in Go
- http://go-database-sql.org how to use SQL databases in Go
- https://dmitri.shuralyov.com/idiomatic-go tips on how to write more idiomatic Go code
- https://divan.github.io/posts/avoid_gotchas will help you avoid gotchas in Go

If you want to learn how to organize your Go project, make sure to read: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp.
Once you are accustomed to the language and syntax, you can read this series of articles for a walkthrough the various standard library packages: https://medium.com/go-walkthrough.

Finally, https://github.com/golang/go/wiki#learning-more-about-go will give a list of even more resources to learn Go`

// GopherNewbieResourcesCommand shows the newbie resources.
var GopherNewbieResourcesCommand = Command{
	Name:        "newbie resources",
	Description: "get a list of newbie resources",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply(gopherNewbieResourcesReply)
	},
}

const gopherRecommendedChannelsReply = `Here is a list of recommended channels:
- https://gophers.slack.com/archives/C2E74QX51 -> talk to the Go team about a certain CL
- https://gophers.slack.com/archives/C02A8LZKT -> for newbie resources
- https://gophers.slack.com/archives/C152YB9UZ -> for remote meetup
- https://gophers.slack.com/archives/C02FKRXAK -> for devops related discussions
- https://gophers.slack.com/archives/C04P1FVLT -> for security related discussions
- https://gophers.slack.com/archives/C02A3DRK6 -> tell the world about the thing you are working on
- https://gophers.slack.com/archives/C0VP8EF3R -> anything and everything performance related
- https://gophers.slack.com/archives/C2VU4UTFZ -> get real time udates from the merged CL for Go itself. For a currated list of important / interesting messages follow: https://twitter.com/golang_cls
- https://gophers.slack.com/archives/C0XDYGUN6 -> Go controlling your bbq grill? Yes, we have that
- https://gophers.slack.com/archives/C029WG6AM -> for code reviews
- https://gophers.slack.com/archives/C08786VLJ -> if you are interested in AWS
- https://gophers.slack.com/archives/C0F1752BB -> for the awesome live podcast
- https://gophers.slack.com/archives/C029WKFFW -> for jobs related to Go`

// GopherRecommendedChannelsCommand shows the recommended channels.
var GopherRecommendedChannelsCommand = Command{
	Name:        "recommended channels",
	Description: "get a list of recommended channels",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply(gopherRecommendedChannelsReply)
	},
}

const gopherDatabaseTutorialReply = `Here's how to work with database/sql in Go: http://go-database-sql.org/`

// GopherDatabaseTutorialCommand shows tuts about sql databases.
var GopherDatabaseTutorialCommand = Command{
	Name:        "database tutorial",
	Description: "tutorial about using sql databases",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply(gopherDatabaseTutorialReply)
	},
}

const gopherPackageLayoutReply = `These articles will explain how to organize your Go packages:
- https://rakyll.org/style-packages/
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp
- https://peter.bourgon.org/go-best-practices-2016/#repository-structure

This article will help you understand the design philosophy for packages: https://www.goinggo.net/2017/02/design-philosophy-on-packaging.html`

// GopherPackageLayoutCommand shows how to structure your Go package.
var GopherPackageLayoutCommand = Command{
	Name:        "package layout",
	Description: "learn how to structure your Go package",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		c.Reply(gopherDatabaseTutorialReply)
	},
}

var gopherLibraryForDynamicReply = func(searchTerm string) string {
	return fmt.Sprintf("You can try to look here: https://godoc.org/?q=%s or here http://go-search.org/search?q=%s", searchTerm, searchTerm)
}

// GopherLibraryForCommand search a go package that matches <name>.
var GopherLibraryForCommand = Command{
	Name:        "library for <name>",
	Description: "learn how to structure your Go package",
	Handler: func(bot *Slack, c hanu.ConversationInterface) {
		name, err := c.String("name")
		if err != nil {
			c.Reply("miss <name> ?")
			return
		}
		c.Reply(gopherLibraryForDynamicReply(name))
	},
}
