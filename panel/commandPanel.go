package panel

import (
	"github.com/gdamore/tcell/v2"

	"ldocker/docker"
)

type CommandPanel struct {
	*BasePanel
	commands        []docker.ListCommand
	onCommandChange func(command docker.ListCommand)
	selectedStyle   tcell.Style
}

func NewCommandPanel(onCommandChange func(command docker.ListCommand)) *CommandPanel {
	commands := []docker.ListCommand{
		{T: "container", F: "-a"},
		{T: "image", F: "-a"},
		{T: "network"},
	}

	return &CommandPanel{
		BasePanel:       NewBasePanel(),
		commands:        commands,
		onCommandChange: onCommandChange,
		selectedStyle:   tcell.StyleDefault.Background(tcell.ColorGreen),
	}
}

func (c *CommandPanel) Draw(screen tcell.Screen) {
	for i, cmd := range c.commands {
		y := c.y + i
		style := c.style

		if i == c.activeLine {
			style = tcell.StyleDefault.Background(tcell.ColorDarkGray)
			if c.focused {
				style = c.selectedStyle
			}
		}

		c.BasePanel.DrawText(screen, y, style, cmd.Type())
	}
}

func (c *CommandPanel) HandleKey(ev *tcell.EventKey) bool {
	handled := c.BasePanel.HandleKey(ev, len(c.commands))

	if handled {
		c.onCommandChange(c.commands[c.activeLine])
	}

	return handled
}

func (c *CommandPanel) GetSelectedCommand() docker.ListCommand {
	return c.commands[c.activeLine]
}
