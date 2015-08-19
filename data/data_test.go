package data

import "testing"
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
