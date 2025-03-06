package panel

import (
	"cmp"
	"slices"

	"github.com/gdamore/tcell/v2"
)

type ResourcePanel struct {
	*BasePanel
	resources        []string
	lines            []string
	onResourceChange func(resource string)
	selectedStyle    tcell.Style
}

func NewResourcePanel(x, y, width int, onResourceChange func(resource string)) *ResourcePanel {
	return &ResourcePanel{
		BasePanel:        NewBasePanel(x, y, width),
		resources:        []string{},
		lines:            []string{},
		onResourceChange: onResourceChange,
		selectedStyle:    tcell.StyleDefault.Background(tcell.ColorGreen),
	}
}

func (r *ResourcePanel) UpdateContent(resources []string, lines []string) {
	r.resources = resources
	r.lines = lines

	if r.activeLine >= len(lines) {
		r.activeLine = 0
	}

	r.maxWidth = len(
		slices.MaxFunc(
			lines, func(a, b string) int {
				return cmp.Compare(len(a), len(b))
			},
		),
	)

	r.onResourceChange(resources[r.activeLine])
}

func (r *ResourcePanel) Draw(screen tcell.Screen) {
	for i, line := range r.lines {

		y := r.y + i + 2
		style := r.style

		if i-1 == r.activeLine {
			style = tcell.StyleDefault.Background(tcell.ColorDarkGray)
			if r.focused {
				style = r.selectedStyle
			}
		}

		r.BasePanel.DrawText(screen, r.x, y, style, line)
	}
}

func (r *ResourcePanel) HandleKey(ev *tcell.EventKey) bool {
	handled := r.BasePanel.HandleKey(ev, len(r.resources))

	if handled {
		r.onResourceChange(r.resources[r.activeLine])
	}

	return handled
}
func (r *ResourcePanel) GetSelectedResource() string {
	if len(r.resources) == 0 {
		return ""
	}
	return r.resources[r.activeLine]
}
