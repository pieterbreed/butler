package todo

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func Test_CanCreateTodo(t *testing.T) {
	var todo *Todo = New("this must get done")
	assert.NotNil(t, todo)
}

func Test_TodoIsATask(t *testing.T) {
	var task Task = New("this must get done")
	assert.False(t, task.IsDone())
	assert.Equal(t, "this must get done", task.What())
}

func Test_CanMarkTodoAsDone(t *testing.T) {
	var task *Todo = New("this must get done")
	assert.NotNil(t, task)
	assert.False(t, task.IsDone())

	var done *Done = task.MarkAsDone()
	assert.NotNil(t, done)
	assert.True(t, done.IsDone())

	var doneTask Task = done
	assert.NotNil(t, doneTask)
	assert.True(t, doneTask.IsDone())
	assert.True(t, strings.Contains(doneTask.What(), "this must get done"))

	// leaves original intact
	assert.NotNil(t, task, "must leave original intact")
	assert.False(t, task.IsDone(), "must leave original intact")
}

func Test_CanMarkAsDoneWithNote(t *testing.T) {
	todo := New("this must get done")
	done := todo.MarkAsDoneWithNote("completion note")
	assert.True(t, strings.Contains(done.What(), "completion note"))

	var doneTask Task = done

	assert.True(t, doneTask.IsDone())
	assert.True(t, strings.Contains(doneTask.What(), "this must get done"))
}

// func Test_MarkingAsDoneCreatesCopy(t *testing.T) {
// 	t.Skip()
// 	var toBeDone Todo = New("this must get done")
// 	assert.False(t, toBeDone.IsDone())

// 	var done Todo = toBeDone.DoneWithNote("completion note")
// 	assert.True(t, done.IsDone())
// 	assert.False(t, toBeDone.IsDone())
// }
