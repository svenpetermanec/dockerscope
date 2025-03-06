package panel

import (
	"github.com/gdamore/tcell/v2"
)

type InspectPanel struct {
	*BasePanel
	lines      []string
	focusStyle tcell.Style
}

func NewInspectPanel(x, y, width int) *InspectPanel {
	return &InspectPanel{
		BasePanel:  NewBasePanel(x, y, width),
		lines:      []string{},
		focusStyle: tcell.StyleDefault.Background(tcell.ColorDarkBlue),
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
		if j < i.activeLine || j >= i.activeLine+i.maxHeight {
			continue
		}

		y := i.y + (j - i.activeLine) + 2

		i.BasePanel.DrawText(screen, i.x+2, y, i.style, line)
	}
}

func (i *InspectPanel) HandleKey(ev *tcell.EventKey) bool {
	return i.BasePanel.HandleKey(ev, len(i.lines))
}
