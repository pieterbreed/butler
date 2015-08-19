package data

import (
	"regexp"
	"testing"
)
import "time"
import "github.com/stretchr/testify/assert"

func Test_CanCreateDotoItem(t *testing.T) {
	_ = NewTodo("testing")
}

func Test_CanMakeCalendarItem(t *testing.T) {
	_ = NewCalendarItem("friends", "V&A", "having dinner", time.Now(), time.Minute)
}

func Test_CanMakeCalenderEntriesImportant(t *testing.T) {
	_ = MakeImportant(NewTodo("testing"))
	_ = MakeImportant(NewCalendarItem("friends", "V&A", "Having dinner", time.Now(), 60*time.Minute))
}

func Test_CanCreateADiary(t *testing.T) {
	_ = MakeDiary()
	p := MakeDiary(
		NewTodo("testing"),
		MakeImportant(NewTodo("testing 2")),
		NewCalendarItem("friends", "V&A", "Having dinner", time.Now(), 60*time.Minute))
	assert.Equal(t, 3, p.Length())
}

func Test_CanModifyADiaryByAddAnItem(t *testing.T) {
	p := MakeDiary()
	assert.Equal(t, 0, p.Length())

	p2 := p.Add(NewTodo("testing"))
	assert.Equal(t, 0, p.Length())
	assert.Equal(t, 1, p2.Length())

	p3 := p2.Add(NewTodo("testing"), MakeImportant(NewTodo("testing 2")))
	assert.Equal(t, 0, p.Length())
	assert.Equal(t, 1, p2.Length())
	assert.Equal(t, 3, p3.Length())
}

func Test_ImportantMakerWorks(t *testing.T) {
	i := ImportantChanger()
	item := i.Change(NewTodo("testing"))
	assert.True(t, IsImportant(item))
	assert.False(t, IsImportant(NewTodo("testing")))
}

func Test_RegexFinderWorks(t *testing.T) {
	rf := RegexOnSummaryFinder(regexp.MustCompile("match"))
	a := NewTodo("testing")
	b := NewTodo("must_match")
	assert.False(t, rf.Finds(a))
	assert.True(t, rf.Finds(b))
}

func Test_ExistingItemsCanBeFoundAndModified(t *testing.T) {
	checker := func(d Diary) int {
		importants := 0
		for i := 0; i < d.Length(); i++ {
			if IsImportant(d.UnsafeGet(i)) {
				importants++
			}
		}
		return importants
	}

	p := MakeDiary(NewTodo("testing1"), NewTodo("testing2"))

	assert.NotNil(t, p)
	assert.Equal(t, 2, p.Length())
	assert.Equal(t, 0, checker(p))

	p2, _ := p.FindAllToChange(
		RegexOnSummaryFinder(regexp.MustCompile("2")),
		ImportantChanger())
	assert.NotNil(t, p2)
	assert.Equal(t, 2, p2.Length())
	assert.Equal(t, 1, checker(p2))
}
