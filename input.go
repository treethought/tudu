package tudu

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	input textinput.Model
}

func NewInput() Input {
	return Input{
		input: textinput.NewModel(),
	}
}
func (m Input) Focus() {
	m.input.Focus()
}

func (m Input) Value() string {
	return m.input.Value()
}

func (m Input) Init() tea.Cmd {
	return textinput.Blink
}

func (m Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.input.Focus()
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, cmdSaveTask(m.Value())
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Input) View() string {
	return m.input.View()
}
