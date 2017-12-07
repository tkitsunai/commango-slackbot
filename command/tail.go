package command

import (
	"fmt"

	"os/exec"

	"github.com/tkitsunai/commango-slackbot/data"
)

type TailCommand struct {
	arg string
	row string
}

func NewTailCommand(row string, arg string) *TailCommand {
	return &TailCommand{
		arg: arg,
		row: row,
	}
}

func (t *TailCommand) Execute() *data.Stdout {
	fmt.Println(t.row)
	out, _ := exec.Command("/bin/sh", "-c", "tail -n"+t.row+" "+t.arg).CombinedOutput()
	fmt.Println(string(out))
	return data.NewStdout(out)
}
