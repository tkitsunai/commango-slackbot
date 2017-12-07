package command

import "github.com/tkitsunai/commango-slackbot/data"

type HelpCommand struct {
	arg string
}

func NewHelpCommand(arg string) *HelpCommand {
	return &HelpCommand{
		arg: arg,
	}
}

func (l *HelpCommand) Execute() *data.Stdout {
	return data.NewStdoutRaw(l.arg + `
Usage:
	ls [FILE_PATH]
		- Using this command, will use the following options by default
			- "-althr"
	tail [LINE COUNT] [FILE]
		- unsupport follow option
`)
}
