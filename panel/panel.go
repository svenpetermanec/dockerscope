package panel

import (
	"github.com/gdamore/tcell/v2"
)

type Panel interface {
	Draw(screen tcell.Screen)
	HandleKey(ev *tcell.EventKey) bool
	Focus()
	Unfocus()
	Resize(x, y, w, h int)
}

type BasePanel struct {
	x      int
	y      int
	width  int
	height int

	focused    bool
	activeLine int
	style      tcell.Style
	horizontal int

	maxWidth int
}

func NewBasePanel() *BasePanel {
	return &BasePanel{
		x:          0,
		y:          0,
		width:      0,
		height:     0,
		focused:    false,
		activeLine: 0,
		style:      tcell.StyleDefault,
		horizontal: 0,
	}
}

func (b *BasePanel) DrawText(screen tcell.Screen, y int, style tcell.Style, line string) {
	x := b.x
	maxWidth := x + b.width

	for _, r := range []rune(line[b.horizontal:]) {
		if x >= maxWidth {
			continue
		}
		screen.SetContent(x, y, r, nil, style)
		x++
	}
}

func (b *BasePanel) HandleKey(ev *tcell.EventKey, maxLines int) bool {
	if !b.focused {
		return false
	}

	switch ev.Key() {
	case tcell.KeyUp:
		if b.activeLine > 0 {
			b.activeLine--
		}
		return true
	case tcell.KeyDown:
		if b.activeLine < maxLines-1 {
			b.activeLine++
		}
		return true
	case tcell.KeyRight:
		if b.horizontal < b.maxWidth-b.width {
			b.horizontal += 1
		}
		return true
	case tcell.KeyLeft:
		if b.horizontal > 1 {
			b.horizontal -= 1
		}
		return true
	default:
		return false
	}
}

func (b *BasePanel) Focus() {
	b.focused = true
}

func (b *BasePanel) Unfocus() {
	b.focused = false
}

func (b *BasePanel) Resize(x, y, w, h int) {
	b.x = x
	b.y = y
	b.width = w
	b.height = h
}
