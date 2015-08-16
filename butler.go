package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	DONE string = "\xe2\x9c\x93"
	NOT_DONE = "\xE2\x9B\xB6"
	CANCELLED string = "\xe2\x9c\x97"
	IMPORTANT string = "!"
	NOTE string = "\xE2\x9C\xB6"
	EYE = "\xE2\x98\xAF"
	APPOINTMENT = "\xE2\x97\x8B"
	FLAG = "\xE2\x9A\x91"
)

var (
	Green = color.New(color.FgGreen, color.Bold).SprintFunc()
	Yellow = color.New(color.FgYellow, color.Bold).SprintFunc()
	Red = color.New(color.FgRed, color.Bold).SprintFunc()
	Gray = color.New(color.FgBlack, color.Bold).SprintFunc()
)

type DiaryItem interface {
	Symbol() string
	Text() string
}
////////////////////////////////////////

type CancelledItem struct {
	item DiaryItem
}
func (di CancelledItem) Symbol() string { return Red(CANCELLED) }
func (di CancelledItem) Text() string { return Gray(di.item.Text()) }
func Cancel(item DiaryItem) CancelledItem { return CancelledItem{item} }

////////////////////////////////////////

type DoneItem struct {
	todo TodoItem
}
func (di DoneItem) Symbol() string { return Green(DONE) }
func (di DoneItem) Text() string { return Gray(di.todo.Text()) }
////////////////////////////////////////

type TodoItem struct {
	task string
}
func (td TodoItem) Symbol() string { return Yellow(FLAG) }
func (td TodoItem) Text() string { return td.task } 
func (td TodoItem) MarkAsDone() DoneItem { return DoneItem{td} }

////////////////////////////////////////

type CalendarItem struct {
	who, where string
	when time.Time 
	duration time.Duration
}
func (ci CalendarItem) NotStarted() bool {
	return ci.when.After(time.Now())
}
func (ci CalendarItem) Busy() bool {
	return ci.when.Before(time.Now()) && ci.when.Add(ci.duration).After(time.Now())
}
func (ci CalendarItem) Done() bool {
	return ci.when.Add(ci.duration).Add(ci.duration).Before(time.Now())
}
func (ci CalendarItem) Symbol() string { 
	switch {
	case ci.NotStarted():
		return Yellow(APPOINTMENT)
	case ci.Busy():
		return Red(APPOINTMENT)
	default:
		return APPOINTMENT
	}
}
func (ci CalendarItem) Text() string {
	switch {
	case ci.NotStarted():
		return fmt.Sprintf("I'm seeing %s at %s, from %s to %s", ci.who, ci.where, ci.when, ci.when.Add(ci.duration))
	case ci.Busy():
		return fmt.Sprintf("I should be seeing %s now at %s until %s", ci.who, ci.where, ci.when.Add(ci.duration))
	default:
		return fmt.Sprintf("I saw %s today at %s", ci.who, ci.where)
	}
}

////////////////////////////////////////

type ImportantItem struct {
	item DiaryItem
}
func (ii ImportantItem) Symbol() string { return fmt.Sprintf("%s%s", Yellow(IMPORTANT), ii.item.Symbol()) }
func (ii ImportantItem) Text() string { return ii.item.Text() }

////////////////////////////////////////

func PrintItem(item DiaryItem) string {
	return fmt.Sprintf("%s %s", item.Symbol(), item.Text())
}

func main() {
	items := make([]DiaryItem, 0)
	items = append(items,
		TodoItem{"get this code written"},
		Cancel(TodoItem{"get this code written"}),
		TodoItem{"perform coding magic"}.MarkAsDone(),
		ImportantItem{TodoItem{"get some sleep"}},
		CalendarItem{"Ferdl", "Somerset Wes", time.Now().Add(-90 * time.Minute), time.Hour},
		Cancel(CalendarItem{"Ferdl", "Somerset Wes", time.Now().Add(-90 * time.Minute), time.Hour}))
	for _, i := range items {
		fmt.Println(PrintItem(i))
	}
}
