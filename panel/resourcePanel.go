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

func NewResourcePanel(onResourceChange func(resource string)) *ResourcePanel {
	return &ResourcePanel{
		BasePanel:        NewBasePanel(),
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

	r.horizontal = 0

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
	visible := r.height - 2
	start := 0

	if r.activeLine > visible {
		start = r.activeLine - visible
	}

	for i := 0; i < r.height && start+i < len(r.lines); i++ {

		y := r.y + i
		style := r.style

		if start+i == r.activeLine+1 {
			style = tcell.StyleDefault.Background(tcell.ColorDarkGray)
			if r.focused {
				style = r.selectedStyle
			}
		}

		r.BasePanel.DrawText(screen, y, style, r.lines[start+i])
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
