package tasklist

import (
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
