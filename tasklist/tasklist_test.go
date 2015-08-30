package tasklist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pieterbreed/butler/todo"
	"github.com/stretchr/testify/assert"
)

func Test_MustCanCreateTaskList(t *testing.T) {
	var _ TaskList = New()
	_ = New(todo.New("this must get done"),
		todo.New("this too"),
		todo.New("maybe this one").Cancel("i don't think so"))
}

func Test_CreatedTaskListIsNotNil(t *testing.T) {
	assert.NotNil(t, New())
	assert.NotNil(t, New(todo.New("this must get done"),
		todo.New("this too"),
		todo.New("maybe this one").Cancel("i don't think so")))
}

func Test_TaskListHasALength(t *testing.T) {
	tl := New()
	assert.Equal(t, 0, tl.Length())
}

func Test_CanAddItemsToATaskListAfterItIsCreated(t *testing.T) {
	t1 := New()
	assert.Equal(t, 0, t1.Length())

	t2 := t1.Add(todo.New("this must get done"))
	assert.Equal(t, 1, t2.Length(), "Needs to actualy have more stuff after you add something")
	assert.Equal(t, 0, t1.Length())
	assert.True(t, reflect.DeepEqual(todo.New("this must get done"), t2.Get(0)))
}

func Test_CanGetItems(t *testing.T) {
	var items []todo.Task = make([]todo.Task, 10)
	for i := 0; i < 10; i++ {
		items[i] = todo.New(fmt.Sprintf("%d", i))
	}

	tl := New(items...)

	for i := 0; i < 10; i++ {
		assert.True(t, reflect.DeepEqual(items[i], tl.Get(i)))
	}
}
