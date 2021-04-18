package tudu

import (
	"github.com/charmbracelet/lipgloss"
	todo "github.com/treethought/todotxt"
	tt "github.com/treethought/todotxt"
)

type TaskView struct {
	todo.Task
}

func NewTaskView(task tt.Task) TaskView {
    return TaskView{task}

}

func (t TaskView) View() string {
	completedColor := lipgloss.Color("#3C3C3C") // a dark gray
	defaultColor := lipgloss.Color("#04B575")   // a green

	if t.Completed {
		style := lipgloss.NewStyle().Background(completedColor)
		return style.Render(t.String())
	}
	style := lipgloss.NewStyle().Background(defaultColor)
	return style.Render(t.String())

}
