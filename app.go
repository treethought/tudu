package tudu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/treethought/boba"
	tt "github.com/treethought/todotxt"
)

type App struct {
	Boba  *boba.App
	tasks *TaskListView
	input *boba.Input
}

func NewApp() *App {
	bapp := boba.NewApp()

	app := &App{Boba: bapp}

	app.tasks = NewTaskListView()
	app.input = boba.NewInput(cmdSaveTask)

	app.Boba.Add("tasks", app.tasks)
	app.Boba.Add("input", app.input)

	app.Boba.SetFocus("tasks")
	app.Boba.SetDelgate(app.delegate)

	return app
}

func (a *App) delegate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tt.TaskList:
		return a.tasks.Update(msg)

	case MessageSaveTask:
		return a.Boba, a.tasks.addTask(msg.taskValue)
	}
	return a.Boba, nil

}
