package tudu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/treethought/boba"
	tt "github.com/treethought/todotxt"
)

type App struct {
	Boba    *boba.App
	tasks   *TaskListView
	input   *boba.Input
	console *Console
	detail  *TaskDetail
}

func NewApp() *App {
	bapp := boba.NewApp()

	app := &App{Boba: bapp}

	app.input = boba.NewInput(cmdSaveTask)
	app.tasks = NewTaskListView()
	app.detail = NewTaskDetail()

	app.Boba.Add("tasks", app.tasks)
	app.Boba.Add("input", app.input)
	app.Boba.Add("detail", app.detail)

	app.Boba.SetFocus("tasks")
	app.Boba.SetDelgate(app.delegate)

	return app
}

func (a *App) delegate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tt.TaskList:
		return a.tasks.Update(msg)

	case MessageViewTask:

		a.detail.SetTask(msg.task)
		return a.Boba, boba.ChangeState("detail")

	case MessageSaveTask:
		return a.Boba, a.tasks.addTask(msg.taskValue)
	}

	return a.Boba, nil

}
