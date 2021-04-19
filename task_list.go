package tudu

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/treethought/boba"
	tt "github.com/treethought/todotxt"
)

type TaskListView struct {
	*boba.List
	tasks      tt.TaskList
	taskMap    map[int]int // mapping of list idx to task id
	filter     *boba.Input
	showFilter bool
}

func NewTaskListView() *TaskListView {
	m := &TaskListView{
		List:    boba.NewList(),
		tasks:   tt.NewTaskList(),
		taskMap: make(map[int]int),
	}

	filter := boba.NewInput(m.filterTasks)
	m.filter = filter

	return m
}

func (m TaskListView) loadTasks() tea.Msg {
	tt.IgnoreComments = false

	tasklist, err := tt.LoadFromFilename("todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	tasklist.Sort(tt.SORT_PRIORITY_ASC)

	return tasklist
}

func (m TaskListView) sortTasks() tea.Cmd {
	return func() tea.Msg {
		tl := m.tasks
		tl.Sort(tt.SORT_CREATED_DATE_DESC)
		return tl
	}
}

func (m *TaskListView) toggleTask() tea.Cmd {
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
		return m.loadTasks()
	}
}

func (m *TaskListView) addTask(value string) tea.Cmd {
	return func() tea.Msg {
		task, err := tt.ParseTask(value)
		if err != nil {
			return nil
		}
		m.tasks.AddTask(task)
		m.tasks.WriteToFilename("todo.txt")
		return m.loadTasks()
	}
}

func (m *TaskListView) filterTasks(query string) tea.Cmd {
    m.showFilter = false
	return func() tea.Msg {
		filtered := filterByString(m.tasks, query)
		return filtered
	}
}

func (m TaskListView) Init() tea.Cmd {
	return m.loadTasks
}

func (m *TaskListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tt.TaskList:
		m.Clear()
		m.tasks = msg
		for _, t := range m.tasks {
			tv := NewTaskView(t)
			m.AddItem(tv)
		}
        m.filter.Blur()
		return m, boba.ChangeState("tasks")

	case tea.KeyMsg:
		if m.showFilter {
			_, cmd = m.filter.Update(msg)
			return m, cmd
		}

		switch msg.String() {

		// search tasks for query
		case "/":
			m.showFilter = true
			return m, nil

		// add task, send state message
		case "a":
			return m, boba.ChangeState("input")

		// sort tasks
		case "s":
			return m, m.sortTasks()

		// toggle completed
		case "x":
			return m, m.toggleTask()
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
	if m.showFilter {
		s += m.filter.View()
	}
	return s
}
