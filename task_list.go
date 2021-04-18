package tudu

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	tt "github.com/treethought/todotxt"
)

type TaskListView struct {
	*List
	tasks   tt.TaskList
	taskMap map[int]int // mapping of list idx to task id
}

func NewTaskListView() TaskListView {
	return TaskListView{
		List:    NewList(),
		tasks:   tt.NewTaskList(),
		taskMap: make(map[int]int),
	}
}

func loadTasks() tea.Msg {
	tt.IgnoreComments = false

	tasklist, err := tt.LoadFromFilename("todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	tasklist.Sort(tt.SORT_PRIORITY_DESC)

	return tasklist
}

func sortTasks(tasks tt.TaskList) tea.Cmd {
	return func() tea.Msg {
		tl := &tasks
		tl.Sort(tt.SORT_PRIORITY_ASC)
		return *tl
	}
}

func toggleTask(m TaskListView) tea.Cmd {
	return func() tea.Msg {
		li := m.CurrentItem()
		tv, ok := li.(TaskView)
		if !ok {
			log.Fatal("not a task")
		}

		task, _ := m.tasks.GetTask(tv.Id)
		if task.Completed {
			task.Reopen()
		} else {
			task.Complete()
		}
		m.tasks.WriteToFilename("todo.txt")
		return loadTasks()
	}
}

func addTask(m TaskListView, value string) tea.Cmd {
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

func (m TaskListView) Init() tea.Cmd {
	return loadTasks
}

func (m TaskListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tt.TaskList:
		m.Clear()
		m.tasks = msg
		for _, t := range m.tasks {
			tv := NewTaskView(t)
			m.AddItem(tv, t.Id)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// add task, send state message
		case "a":
            return m, cmdChangeState(stateInput)

		// sort tasks
		case "s":
			return m, sortTasks(m.tasks)

		// toggle completed
		case "x":
			return m, toggleTask(m)
		}
		_, cmd := m.List.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m TaskListView) View() string {
	// The header
	s := "task list:\n\n"

	s += m.List.View()
	return s
}
