package data

type Stdout struct {
	out string
}

func NewStdoutRaw(out string) *Stdout {
	return &Stdout{
		out: out,
	}
}

func NewStdout(b []byte) *Stdout {
	return &Stdout{
		out: string(b),
	}
}

func (s *Stdout) String() string {
	return "```" + s.out + "```"
}
