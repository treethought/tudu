package tudu

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	todo "github.com/treethought/todotxt"
	tt "github.com/treethought/todotxt"
)

var (
	enableStyle    = true
	completedColor = lipgloss.Color("#3C3C3C") // a dark gray
	projectColor   = lipgloss.Color("#04B575") // a green
	contextColor   = lipgloss.Color("#e28743") // a green
)

type TaskView struct {
	todo.Task
}

func NewTaskView(task tt.Task) TaskView {
	return TaskView{task}

}

func (t TaskView) styledView() string {
	style := lipgloss.NewStyle()

	if t.Completed {
		style = style.Background(completedColor).Strikethrough(true)
        return style.Render(t.String())
	} 


	s := t.String()

	for _, p := range t.Projects {
		pstyle := style.Foreground(projectColor)
		s = strings.ReplaceAll(s, p, pstyle.Render(p))
	}

	for _, c := range t.Contexts {
		cstyle := style.Foreground(contextColor)
		s = strings.ReplaceAll(s, c, cstyle.Render(c))
	}

	return s

}

func (t TaskView) View() string {
	if enableStyle {
		return t.styledView()
	}

	return t.String()

}
