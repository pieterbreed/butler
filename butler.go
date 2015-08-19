package main

import (
	"fmt"
	"time"
	"bytes"
	"github.com/fatih/color"
	"./data"
)

const (
	DONE string = "\xe2\x9c\x93"
//	NOT_DONE = "\xE2\x9B\xB6"
	NOT_DONE = "x"
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

func RenderSymbol(item data.DiaryItem) string {
	switch t := item.(type) {
	case *data.TodoItem:
		return Green(NOT_DONE)
	case *data.ImportantItem:
		return fmt.Sprintf(
			"%s%s",
			Red(IMPORTANT),
			RenderSymbol(t.Item))
	default: return "-"
	}
}

func RenderItem(i data.DiaryItem) string {
	return fmt.Sprintf("(%s) %s", RenderSymbol(i), i.Summary())
}

func RenderDiary(d data.Diary) string {
	var result bytes.Buffer
	for i := 0; i < d.Length(); i++ {
		result.WriteString(RenderItem(d.UnsafeGet(i)))
		result.WriteString("\n")
	}
	return result.String()
}

func main() {
	d := data.MakeDiary(
		data.NewTodo("testing"),
		data.MakeImportant(data.NewTodo("testing2")),
		data.NewCalendarItem("Izak", "Panchos", "dinner", time.Now(), time.Hour))
	fmt.Print(RenderDiary(d))
}
