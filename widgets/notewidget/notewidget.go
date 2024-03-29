package notewidget

import (
	"context"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type NoteObject struct {
	Title string
	Body  string
	Date  time.Time
}

func MakeNote(title, body string) NoteObject {
	return NoteObject{
		Title: title,
		Body:  body,
		Date:  time.Now(),
	}
}

type NoteUser struct {
	lock    sync.Mutex
	Notes   []NoteObject
	Context *context.Context
}

func MakeNoteUser(ctx *context.Context) *NoteUser {
	note1 := MakeNote("My first note", "This is my first note body")
	note2 := MakeNote("My seconds note", "This is my second note body")
	note3 := MakeNote("My third note", "This is my third note body")
	noteUser := NoteUser{
		Notes:   []NoteObject{note1, note2, note3},
		Context: ctx,
	}

	return &noteUser
}

func (nu *NoteUser) CreateNote(title, body string) {
	nu.lock.Lock()
	defer nu.lock.Unlock()
	newNote := MakeNote(title, body)
	nu.Notes = append(nu.Notes, newNote)
	runtime.EventsEmit(*nu.Context, "new_note", newNote)
	/* 	return newNote */
}

func (nu *NoteUser) startUpdate(index int) func(string, string) {
	nu.lock.Lock()
	defer nu.lock.Unlock()
	return func(title, body string) {
	}
}

func (nu *NoteUser) GetNotes() []NoteObject {
	nu.lock.Lock()
	defer nu.lock.Unlock()
	return nu.Notes
}

// func (nu *NoteUser) updateNote(index int, title string, body string) {

// }
