package tasklist

import (
	"github.com/pieterbreed/butler/todo"
)

type TaskList []todo.Task

func New(items ...todo.Task) TaskList {
	return append(make([]todo.Task, 0, len(items)), items...)
}
