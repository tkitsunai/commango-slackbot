package command

import (
	"fmt"
	"os/exec"

	"github.com/tkitsunai/commango-slackbot/data"
)

type LsCommand struct {
	arg string
}

func NewLsCommand(arg string) *LsCommand {
	return &LsCommand{
		arg: arg,
	}
}

func (l *LsCommand) Execute() *data.Stdout {
	out, _ := exec.Command("/bin/sh", "-c", "ls -althr "+l.arg+" | tail -20").CombinedOutput()
	fmt.Println(string(out))
	return data.NewStdout(out)
}
