package panel

import (
	"github.com/gdamore/tcell/v2"
)

type InspectPanel struct {
	*BasePanel
	lines []string
}

func NewInspectPanel() *InspectPanel {
	return &InspectPanel{
		BasePanel: NewBasePanel(),
		lines:     []string{},
	}
}

func (i *InspectPanel) UpdateContent(lines []string) {
	i.lines = lines

	if i.activeLine >= len(lines) {
		i.activeLine = 0
	}
}

func (i *InspectPanel) Draw(screen tcell.Screen) {
	for j, line := range i.lines {
		if j < i.activeLine || j >= i.activeLine+i.height {
			continue
		}

		y := i.y + (j - i.activeLine)

		i.BasePanel.DrawText(screen, y, i.style, line)
	}
}

func (i *InspectPanel) HandleKey(ev *tcell.EventKey) bool {
	return i.BasePanel.HandleKey(ev, len(i.lines))
}
