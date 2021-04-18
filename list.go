package tudu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Viewer interface {
	View() string
}

type SelectedFunc func(msg tea.Msg) (tea.Model, tea.Cmd)

type List struct {
	items  []Viewer
	cursor int
}

func NewList() *List {
	return &List{
		items:  []Viewer{},
		cursor: 0,
		// itemMap: make(map[int]interface{}),
	}
}

func (m *List) Clear() {
	m.items = nil
	// m.itemMap = make(map[int]interface{})
}

func (m *List) AddItem(item Viewer, ref interface{}) {
	m.items = append(m.items, item)
	// m.itemMap[len(m.items)-1] = ref
}

func (m *List) CurrentItem() Viewer {
	return m.items[m.cursor]
}

func (m *List) Init() tea.Cmd {
	return nil
}

func (m *List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m *List) View() string {
	s := ""
	for idx, item := range m.items {
		cursor := "  "
		if idx == m.cursor {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s\n", cursor, item.View())
	}

	return s
}
