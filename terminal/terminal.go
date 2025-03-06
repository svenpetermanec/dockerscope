package terminal

import (
	"os"

	"github.com/gdamore/tcell/v2"

	"ldocker/docker"
	"ldocker/panel"
)

type Terminal struct {
	screen         tcell.Screen
	panels         []panel.Panel
	activePanel    int
	dockerExecutor *docker.Executor
	defaultStyle   tcell.Style
}

func NewTerminal(dockerExecutor *docker.Executor) (*Terminal, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	err = screen.Init()
	if err != nil {
		panic(err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)

	return &Terminal{
		screen:         screen,
		panels:         make([]panel.Panel, 0),
		activePanel:    0,
		dockerExecutor: dockerExecutor,
		defaultStyle:   defStyle,
	}, nil
}

func (t *Terminal) AddPanel(panel panel.Panel) {
	t.panels = append(t.panels, panel)
	if len(t.panels) == 1 {
		panel.Focus()
	}
}

func (t *Terminal) Start() {
	for {
		t.screen.Clear()
		for _, p := range t.panels {
			p.Draw(t.screen)
		}
		t.screen.Show()

		event := t.screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			handled := t.panels[t.activePanel].HandleKey(ev)
			if !handled {
				t.handleGlobalKey(ev)
			}
		case *tcell.EventResize:
			t.screen.Sync()
		}
	}
}

func (t *Terminal) handleGlobalKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		t.screen.Fini()
		os.Exit(0)

	case tcell.KeyTab:
		t.panels[t.activePanel].Unfocus()

		if t.activePanel == 2 {
			t.activePanel = 0
		} else {
			t.activePanel++
		}

		t.panels[t.activePanel].Focus()

	case tcell.KeyBacktab:
		t.panels[t.activePanel].Unfocus()

		if t.activePanel == 0 {
			t.activePanel = 2
		} else {
			t.activePanel--
		}

		t.panels[t.activePanel].Focus()
	default:
		return
	}
}
