package tasklist

import (
	"github.com/pieterbreed/butler/todo"
)

type TaskList []todo.Task

func New(items ...todo.Task) TaskList {
	return append(make([]todo.Task, 0, len(items)), items...)
}

func (t TaskList) Length() int {
	return len(t)
}

func (t TaskList) Add(items ...todo.Task) TaskList {
	return append(make(
		[]todo.Task,
		0,
		len(t) + len(items)), append(t, items...)...)
}

func (t TaskList) Get(i int) todo.Task {
	return t[i]
}
