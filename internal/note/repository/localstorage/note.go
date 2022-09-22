package localstorage

import (
	"context"
	"sync"

	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/lekan-rust/notes-app/internal/note"
)

type NoteLocalStorage struct {
	notes map[string]*models.Note
	mutex *sync.Mutex
}

func NewNoteLocalStorage() *NoteLocalStorage {
	return &NoteLocalStorage{
		notes: make(map[string]*models.Note),
		mutex: new(sync.Mutex),
	}
}

// CreateNote is creates note in local storage.
func (s *NoteLocalStorage) CreateNote(ctx context.Context, user *models.User, note *models.Note) error {
	note.UserID = user.ID

	s.mutex.Lock()
	s.notes[note.ID] = note
	s.mutex.Unlock()

	return nil
}

// GetNotes takes all notes from local storage by some user.
func (s *NoteLocalStorage) GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error) {
	notes := make([]*models.Note, 0)

	s.mutex.Lock()
	for _, note := range s.notes {
		if note.UserID == user.ID {
			notes = append(notes, note)
		}
	}
	s.mutex.Unlock()

	return notes, nil
}

// DeleteNote removes note from local storage.
func (s *NoteLocalStorage) DeleteNote(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	n, ok := s.notes[id]
	if ok && n.UserID == user.ID {
		delete(s.notes, id)
		return nil
	}
	return note.ErrNoteNotFound
}
