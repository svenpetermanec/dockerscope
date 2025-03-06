package docker

type Command interface {
	Type() string
	Action() string
	Flag() string
}

type ListCommand struct {
	T string
	F string
}

func (l ListCommand) Type() string {
	return l.T
}

func (l ListCommand) Action() string {
	return "ls"
}

func (l ListCommand) Flag() string {
	return l.F
}

type InspectCommand struct {
	T string
	R string
}

func (i InspectCommand) Type() string {
	return i.T
}

func (i InspectCommand) Action() string {
	return "inspect"
}

func (i InspectCommand) Flag() string {
	return i.R
}
