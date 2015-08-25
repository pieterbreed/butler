package todo

import (
	"fmt"
)

type Task interface {
	IsDone() bool
	What() string
}

func New(what string) *Todo { return &Todo{what} }

type Todo struct {
	what string
}

func (*Todo) IsDone() bool { return false }
func (t *Todo) What() string { return t.what }

func (t *Todo) MarkAsDone() *Done {
	return &Done{t}
}

type Done struct {
	todo *Todo
}

func (*Done) IsDone() bool { return true }
func (t *Done) What() string { return t.todo.What() }

func (t *Todo) MarkAsDoneWithNote(note string) *DoneWithNote {
	return &DoneWithNote{t, note}
}

type DoneWithNote struct {
	todo *Todo
	note string
}

func (*DoneWithNote) IsDone() bool { return true }
func (n *DoneWithNote) What() string { 
	return fmt.Sprintf("DONE: %s (%s)", n.note, n.todo.What())
}
