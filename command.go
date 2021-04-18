package tudu

import tea "github.com/charmbracelet/bubbletea"

func cmdChangeState(newState state) tea.Cmd {
	return func() tea.Msg {
		return MessageState{state: newState}
	}
}

func  cmdSaveTask(val string) tea.Cmd {
	return func() tea.Msg {
		return MessageSaveTask{taskValue: val}
	}
}
