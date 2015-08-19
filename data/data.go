package data

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////

// the basic line item in a diary
type DiaryItem interface {
	Summary() string // the summary one-liner for this item
}

////////////////////////////////////////////////////////////

// represents a single atomic task that can be in a note-yet-done state
type TodoItem struct {
	Task string // the task that must be performed
}

func (td *TodoItem) Summary() string {
	return td.Task
}
func NewTodo(task string) *TodoItem {
	return &TodoItem{task}
}

////////////////////////////////////////////////////////////

// a calendar entry, see some -body/-people somewhere at some time for some duration of time
type CalendarItem struct {
	who, where, forwhat string
	when                time.Time
	duration            time.Duration
}

func (ci *CalendarItem) Summary() string {
	return fmt.Sprintf("CalendarItem: %v", *ci)
}
func NewCalendarItem(who, where, forwhat string, when time.Time, duration time.Duration) *CalendarItem {
	return &CalendarItem{who, where, forwhat, when, duration}
}

////////////////////////////////////////////////////////////

// wrapper type, means the wrapped item is important
type ImportantItem struct {
	item DiaryItem
}

func (ii *ImportantItem) Summary() string {
	return ii.item.Summary()
}

func MakeImportant(item DiaryItem) DiaryItem {
	return &ImportantItem{item}
}

////////////////////////////////////////////////////////////

// represents a project of some kind
type Diary interface {
	Length() int // how many items are in this project
}
type DiaryModifier interface {
	Add(items ...DiaryItem)
}

type DiaryItemsDiary struct {
	items []DiaryItem
}

func (d *DiaryItemsDiary) Length() int {
	return len(d.items)
}

func MakeDiary(items ...DiaryItem) *DiaryItemsDiary {
	return &DiaryItemsDiary{append(make([]DiaryItem, 0, len(items)), items...)}
}

func (d *DiaryItemsDiary) Add(items ...DiaryItem) *DiaryItemsDiary {
	return &DiaryItemsDiary{append(d.items, items...)}
}
