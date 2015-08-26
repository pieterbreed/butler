package todo

import (
	"fmt"
)

// all states a task can be in have to actually "be" tasks, ie, they have to satisfy
// this Task interface
type Task interface {
	IsDone() bool      // task is successfully completed
	IsCancelled() bool // task is un-successfully completed
	IsCompleted() bool // task is un-completed
	What() string
}

func New(what string) *Todo { return &Todo{what} }

// ----------------------------------------
// todo - a starting state

type Todo struct {
	what string
}

func (*Todo) IsDone() bool      { return false }
func (*Todo) IsCancelled() bool { return false }
func (*Todo) IsCompleted() bool { return false }
func (t *Todo) What() string    { return t.what }

func (t *Todo) MarkAsDone() *Done {
	return &Done{t}
}
func (t *Todo) MarkAsDoneWithNote(note string) *DoneWithNote {
	return &DoneWithNote{t, note}
}
func (t *Todo) Cancel(note string) *Cancelled {
	return &Cancelled{t, note}
}

// ----------------------------------------
// done

type Done struct {
	todo *Todo
}

func (*Done) IsDone() bool      { return true }
func (*Done) IsCancelled() bool { return false }
func (*Done) IsCompleted() bool { return true }
func (t *Done) What() string    { return t.todo.What() }

// ----------------------------------------
// done with note

type DoneWithNote struct {
	todo *Todo
	note string
}

func (*DoneWithNote) IsDone() bool      { return true }
func (*DoneWithNote) IsCancelled() bool { return false }
func (*DoneWithNote) IsCompleted() bool { return true }
func (n *DoneWithNote) What() string {
	return fmt.Sprintf("%s (%s)", n.note, n.todo.What())
}

// ----------------------------------------
// cancelled with note

type Cancelled struct {
	todo *Todo
	note string
}

func (*Cancelled) IsDone() bool      { return false }
func (*Cancelled) IsCancelled() bool { return true }
func (*Cancelled) IsCompleted() bool { return true }
func (n *Cancelled) What() string {
	return fmt.Sprintf("%s (%s)", n.note, n.todo.What())
}
