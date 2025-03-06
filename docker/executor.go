package docker

import (
	"bufio"
	"os/exec"
	"strings"
)

type Executor struct {
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) ExecuteCommand(cmd Command) ([]string, []string) {
	args := []string{cmd.Type(), cmd.Action()}
	if cmd.Flag() != "" {
		args = append(args, cmd.Flag())
	}
	
	command := exec.Command("docker", args...)
	out, _ := command.StdoutPipe()

	err := command.Start()
	if err != nil {
		panic(err)
	}

	resources := make([]string, 0)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		line := scanner.Text()
		resources = append(resources, strings.SplitN(line, " ", 2)[0])

		lines = append(lines, line)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return resources[1:], lines
}
