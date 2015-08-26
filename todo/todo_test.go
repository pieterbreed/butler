package todo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateTodo(t *testing.T) {
	var todo *Todo = New("this must get done")
	assert.NotNil(t, todo)
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

func Test_AllStatesAreTasks(t *testing.T) {
	todo := New("this must get done")
	var _ Task = todo
	var _ Task = (todo.MarkAsDone())
	var _ Task = (todo.MarkAsDoneWithNote("note"))
	var _ Task = (todo.Cancel("cancelation note"))
}

func Tast_CanCancelTask(t *testing.T) {
	todo := New("this must get done")
	cancelled := todo.Cancel("note")
	assert.True(t, cancelled.IsCancelled())
	assert.True(t, cancelled.IsCompleted())
	assert.False(t, cancelled.IsDone())
}
