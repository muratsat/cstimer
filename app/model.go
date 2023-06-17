package app

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/muratsat/cstimer/app/cube"
)

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

type model struct {
	stopwatch stopwatch.Model
	cube      cube.Cube3x3
	help      help.Model
	keymap    keymap
	moves     []cube.Move
}

func New() model {
	return model{
		stopwatch: stopwatch.NewWithInterval(time.Millisecond * 30),
		moves:     cube.GenerateMoves(),
		cube:      cube.NewCube3x3(),
		help:      help.New(),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("space", "s"),
				key.WithHelp("s", "start/stop")),
			stop: key.NewBinding(
				key.WithKeys("s", " "),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset")),
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "q"),
				key.WithHelp("q", "quit")),
		},
	}
}

func (m model) Init() tea.Cmd {
	return m.stopwatch.Reset()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case " ", "s":
			if m.stopwatch.Running() {
				return m, m.stopwatch.Stop()
			} else {
				return m, m.stopwatch.Start()
			}
		case "r":
			m.moves = cube.GenerateMoves()
			return m, m.stopwatch.Reset()
		}
	}

	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m model) HelpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.start,
		m.keymap.stop,
		m.keymap.reset,
		m.keymap.quit,
	})
}

func (m model) View() string {
	scramble := cube.ScrambleString(m.moves)
	result := "scramble:   " + scramble + "\n\n\n"
	// result += m.cube.String() + "\n\n"
	result += "   time: " + m.stopwatch.View() + "\n\n"
	result += m.HelpView()
	return result
}
