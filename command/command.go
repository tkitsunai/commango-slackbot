package command

import "github.com/tkitsunai/commango-slackbot/data"

type Command interface {
	Execute() *data.Stdout
}
