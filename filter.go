package tudu

import (
	"strings"

	tt "github.com/treethought/todotxt"
)

type FilterFunc func(tasks tt.TaskList, query string) tt.TaskList

func filterByProject(tasks tt.TaskList, project string) tt.TaskList {
	taskList := tasks.Filter(func(task tt.Task) bool {
		for _, p := range task.Projects {
			if p == project {
				return true
			}
		}
		return false
	})
	return *taskList
}

func filterByContext(tasks tt.TaskList, context string) tt.TaskList {
	taskList := tasks.Filter(func(task tt.Task) bool {
		for _, c := range task.Contexts {
			if c == context {
				return true
			}
		}
		return false
	})
	return *taskList
}

func filterByString(tasks tt.TaskList, query string) tt.TaskList {
	taskList := tasks.Filter(func(task tt.Task) bool {
		if strings.Contains(task.Todo, query) {
			return true
		}
		return false
	})
	return *taskList
}
