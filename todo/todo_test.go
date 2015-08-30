package todo

import (
	"strings"
	"testing"
	"reflect"

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

func Test_CanCancelTask(t *testing.T) {
	todo := New("this must get done")
	cancelled := todo.Cancel("note")
	assert.True(t, cancelled.IsCancelled())
	assert.True(t, cancelled.IsCompleted())
	assert.False(t, cancelled.IsDone())
}

func Test_EqualsWorksAsExpected(t *testing.T) {
	t1a := New("this must get done")
	t1b := New("this must get done")
	t2 := New("this too")

	ef := func(a, b Task, exp bool) {
		assert.True(t, exp == reflect.DeepEqual(a, b))
	}

	ef(t1a, t1b, true)
	ef(t1a, t2, false)

	ef(t1a.MarkAsDone(), t1b.MarkAsDone(), true)
	ef(t1a.MarkAsDone(), t2.MarkAsDone(), false)

	ef(t1a.MarkAsDoneWithNote("note"), t1b.MarkAsDoneWithNote("note"), true)
	ef(t1a.MarkAsDoneWithNote("note"), t2.MarkAsDoneWithNote("note"), false)

	ef(t1a.Cancel("note"), t1b.Cancel("note"), true)
	ef(t1a.Cancel("note"), t2.Cancel("note"), false)
}
