package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	tt "github.com/treethought/todotxt"
)

type model struct {
	tasks       tt.TaskList
	cursor      int         // which to-do list item our cursor is pointing at
	taskMap     map[int]int // mapping of list idx to task id
	input       textinput.Model
	inputActive bool
}

func loadTasks() tea.Msg {
	tt.IgnoreComments = false

	tasklist, err := tt.LoadFromFilename("todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	tasklist.Sort(tt.SORT_COMPLETED_DATE_ASC)

	return tasklist
}

func sortTasks(tasks tt.TaskList) tea.Cmd {
	return func() tea.Msg {
		tl := &tasks
		tl.Sort(tt.SORT_COMPLETED_DATE_DESC)
		return *tl
	}
}

func toggleTask(m model) tea.Cmd {
	return func() tea.Msg {
		tid, ok := m.taskMap[m.cursor]
		if !ok {
			return nil
		}

		task, _ := m.tasks.GetTask(tid)
		if task.Completed {
			task.Reopen()
		} else {
			task.Complete()
		}
		m.tasks.WriteToFilename("todo.txt")
		return loadTasks()
	}
}

func addTask(m model, value string) tea.Cmd {
	return func() tea.Msg {
		task, err := tt.ParseTask(value)
		if err != nil {
			log.Fatal(err)
		}
		m.tasks.AddTask(task)
		m.tasks.WriteToFilename("todo.txt")
		return loadTasks()
	}
}

var initialModel = model{
	// Our to-do list is just a grocery list
	tasks: []tt.Task{},

	// mapping of list index to task id
	taskMap: make(map[int]int),
	input:   textinput.NewModel(),
}

func (m model) Init() tea.Cmd {
	return loadTasks
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tt.TaskList:
		m.tasks = msg
		for i, t := range m.tasks {
			m.taskMap[i] = t.Id
		}
		return m, nil

	// Is it a key press?
	case tea.KeyMsg:

		if m.inputActive {
			switch msg.Type {
			case tea.KeyCtrlC:
				fallthrough
			case tea.KeyEsc:
				fallthrough
			case tea.KeyEnter:
				m.inputActive = false
				return m, addTask(m, m.input.Value())
			default:
				input, cmd := m.input.Update(msg)
				m.input = input
				return m, cmd
			}
		}

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		// add task
		case "a":
			m.inputActive = true
            m.input.Focus()
			return m, nil

		// sort tasks
		case "s":
			return m, sortTasks(m.tasks)

		// toggle completed
		case "x":
			return m, toggleTask(m)

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "task list:\n\n"

	if m.inputActive {

		return fmt.Sprintf("Add new task:\n%s", m.input.View())
	}

	// Iterate over our choices
	for i, task := range m.tasks {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		// checked := " " // not selected
		// if _, ok := m.selected[i]; ok {
		// 	checked = "x" // selected!
		// }

		// Render the row
		s += fmt.Sprintf("%s  %s\n", cursor, task)
	}

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
