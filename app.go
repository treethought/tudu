package tudu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	tt "github.com/treethought/todotxt"
)

type state int

const (
	stateTasks state = iota
	stateInput
)

type App struct {
	state state
	tasks tea.Model
	input tea.Model
}

func NewApp() App {

	app := App{}

	app.tasks = NewTaskListView()
	app.input = NewInput()
	return app
}

func (m App) Init() tea.Cmd {
	m.state = stateTasks
	return m.tasks.Init()
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case MessageSaveTask:
		m.state = stateTasks
		tasksList, _ := m.tasks.(TaskListView)
		return m, addTask(tasksList, msg.taskValue)

	// state change message to focus new view
	case MessageState:
		m.state = state(msg.state)

		switch m.state {

		case stateInput:
			input := NewInput()
			input.Focus()
			m.input = input
			return m, nil
		}

	// update task list for task chganges
	case tt.TaskList:
		tasks, cmd := m.tasks.Update(msg)
		m.tasks = tasks
		return m, cmd

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	switch m.state {
	case stateTasks:
		taskList, cmd := m.tasks.Update(msg)
		m.tasks = taskList
		return m, cmd

	case stateInput:
		input, cmd := m.input.Update(msg)
		m.input = input
		return m, cmd
	}

	return m, nil
}

func (m App) View() string {
	switch m.state {
	case stateInput:
		return fmt.Sprintf("Add new task:\n%s", m.input.View())
	default:
		return m.tasks.View()

	}
}
