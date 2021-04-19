package tudu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/treethought/boba"
)

var options = []string{"projects", "contexts"}

type Console struct {
	input *boba.Input
	list  *boba.List
}

func NewConsole() *Console {
	m := &Console{}
	m.list = boba.NewList()
	m.input = boba.NewInput(m.selected)
	return m

}
func (m Console) Focus() {
	m.input.Focus()
}

func (m Console) selected(val string) tea.Cmd {
	return nil
}

func (m Console) Value() string {
	return m.input.Value()
}

func (m Console) Init() tea.Cmd {
	for _, o := range options {
		m.list.Add(o)
	}
	return textinput.Blink
}

func (m Console) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.input.Focus()
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, cmdSaveTask(m.Value())
		}
	}

	if m.Value() == "" {
		_, cmd = m.input.Update(msg)
		return m, cmd
	}

	m.list.Clear()
	for _, o := range options {
		if strings.Contains(o, m.Value()) {
			m.list.Add(o)
		}
	}

	_, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Console) View() string {
	return fmt.Sprintf(": %s\n%s", m.input.View(), m.list.View())
}
