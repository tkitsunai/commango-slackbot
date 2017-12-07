package command

import "errors"

func FactoryCommand(command string, flag []string) (Command, error) {
	switch command {
	case Help.String():
		return NewHelpCommand(""), nil
	case Ls.String():
		if len(flag) >= 2 {
			return NewLsCommand(flag[1]), nil
		}
	case Tail.String():
		if len(flag) == 2 {
			return NewTailCommand("30", flag[2]), nil
		}
		if len(flag) >= 3 {
			return NewTailCommand(flag[1], flag[2]), nil
		}
	default:
		return nil, errors.New("unsupported command")
	}
	return NewHelpCommand("illegal option. please readme help"), nil
}
