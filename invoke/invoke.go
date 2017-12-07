package invoke

import (
	"fmt"

	"github.com/tkitsunai/commango-slackbot/command"
	"github.com/tkitsunai/commango-slackbot/data"
)

func CallCommand(commands []string) *data.Stdout {
	if len(commands) < 1 {
		return data.NewStdoutRaw("The number of arguments is not enough")
	}

	command, err := command.FactoryCommand(commands[0], commands)

	if err != nil {
		return data.NewStdoutRaw(fmt.Sprintln("command :", commands[0], err.Error()))
	}

	return command.Execute()
}
