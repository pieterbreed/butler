package main

import (
	"fmt"
	"time"
//	"regexp"

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
	SuperWhite = color.New(color.FgWhite, color.Bold).SprintFunc()
	Pink = color.New(color.FgRed).SprintFunc()
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
func (ii ImportantItem) Symbol() string { return fmt.Sprintf("%s%s", ii.item.Symbol(), Pink(IMPORTANT)) }
func (ii ImportantItem) Text() string { return SuperWhite(ii.item.Text()) }
func MakeImportant(item DiaryItem) DiaryItem {
	importantItem, ok := item.(ImportantItem)
	if ok { return importantItem }
	return ImportantItem{item}
}

////////////////////////////////////////

type Diary []DiaryItem

func NewDiary() Diary {
	return NewDiaryN(0)
}

func NewDiaryN(n int) Diary {
	return make([]DiaryItem, n)
}

func (d Diary) Length() int {
	return len(d)
}

// adds to the end of the diary, like if you were writing on paper
func (d Diary) Add(item ...DiaryItem) Diary {
	return append(d, item...)
}

func (d Diary) Select(selector func(i int, di DiaryItem)) {
	for i, di := range d {
		selector(i, di)
	}
}

func (d Diary) Modify(predicate func (i DiaryItem) bool, modification func (i DiaryItem) DiaryItem) Diary {
	result := NewDiaryN(d.Length())
	for i := 0; i < len(d); i++ {
		if predicate(d[i]) {
			result = append(result, modification(d[i]))
		} else {
			result = append(result, d[i])
		}
	}
	return result
}

////////////////////////////////////////

func PrintItem(item DiaryItem) string {
	return fmt.Sprintf("%s %s", item.Symbol(), item.Text())
}

func SelectorFromRegexp(_ string) (func(DiaryItem) bool) {
	return func(item DiaryItem) bool {
		return true
	}
}

func main() {
	items := NewDiary()
	items = items.Add(TodoItem{"get ready for work"}).Modify(SelectorFromRegexp("aoeu"), MakeImportant)
	

// make([]DiaryItem, 0)
// 	items = append(items,
// 		TodoItem{"get this code written"},
// 		Cancel(TodoItem{"get this code written"}),
// 		TodoItem{"perform coding magic"}.MarkAsDone(),
// 		MakeImportant(TodoItem{"get some sleep"}),
// 		CalendarItem{"Ferdl", "Somerset Wes", time.Now().Add(-90 * time.Minute), time.Hour},
// 		Cancel(CalendarItem{"Ferdl", "Somerset Wes", time.Now().Add(-90 * time.Minute), time.Hour}))
	for _, i := range items {
		fmt.Println(PrintItem(i))
	}
}
