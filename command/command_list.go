package command

type List int

const (
	Help List = iota
	Ls
	Tail
)

func (cl List) String() string {
	switch cl {
	case Help:
		return "help"
	case Ls:
		return "ls"
	case Tail:
		return "tail"
	default:
		return "Unknown"
	}
}
