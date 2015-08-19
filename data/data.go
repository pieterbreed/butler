package data

import (
	"fmt"
	"regexp"
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
	Length() int                                          // how many items are in this project
	Get(i int) (DiaryItem, error)                         // returns the diaryitem at index i
	UnsafeGet(i int) DiaryItem                            // panics if index is out of range
	FindAllToChange(ItemFinder, ItemChanger) (Diary, int) // select items with ItemFinder, change with ItemChanger, get a new Diary back and how many changes were made
}
type ItemChanger interface {
	Change(DiaryItem) DiaryItem
}
type ItemFinder interface {
	Finds(DiaryItem) bool
}

type regexOnSummaryFinder struct {
	regex *regexp.Regexp
}

func (rf regexOnSummaryFinder) Finds(item DiaryItem) bool {
	return rf.regex.MatchString(item.Summary())
}

func RegexOnSummaryFinder(r *regexp.Regexp) ItemFinder {
	return regexOnSummaryFinder{r}
}

type importantMaker struct{}

func (importantMaker) Change(item DiaryItem) DiaryItem {
	isImportant, ok := item.(*ImportantItem)
	if ok {
		return isImportant
	}
	return MakeImportant(item)
}
func ImportantChanger() ItemChanger {
	return importantMaker{}
}
func IsImportant(item DiaryItem) bool {
	_, ok := item.(*ImportantItem)
	return ok
}

type DiaryItemsDiary struct {
	items []DiaryItem
}

func (d *DiaryItemsDiary) Length() int {
	return len(d.items)
}

func (d *DiaryItemsDiary) UnsafeGet(i int) DiaryItem {
	item, err := d.Get(i)
	if err != nil {
		panic(err)
	}
	return item
}

func (d *DiaryItemsDiary) Get(i int) (DiaryItem, error) {
	if i > len(d.items) {
		return nil, fmt.Errorf("Index out of range for Diary")
	}
	return d.items[i], nil
}

func (d *DiaryItemsDiary) FindAllToChange(
	finder ItemFinder,
	changer ItemChanger) (Diary, int) {

	result := &DiaryItemsDiary{make([]DiaryItem, 0, len(d.items))}
	changes := 0

	for i := 0; i < len(d.items); i++ {
		if finder.Finds(d.items[i]) {
			result.items = append(result.items, changer.Change(d.items[i]))
			changes++
		} else {
			result.items = append(result.items, d.items[i])
		}
	}

	return result, changes
}

func MakeDiary(items ...DiaryItem) *DiaryItemsDiary {
	return &DiaryItemsDiary{append(make([]DiaryItem, 0, len(items)), items...)}
}

func (d *DiaryItemsDiary) Add(items ...DiaryItem) *DiaryItemsDiary {
	return &DiaryItemsDiary{append(d.items, items...)}
}
