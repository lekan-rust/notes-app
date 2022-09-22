package localstorage

import (
	"context"
	"fmt"
	"testing"

	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/lekan-rust/notes-app/internal/note"
	"github.com/stretchr/testify/assert"
)

func TestGetNotes(t *testing.T) {
	id := "id"
	user := &models.User{ID: id}

	s := NewNoteLocalStorage()

	for i := 0; i < 10; i++ {
		note := &models.Note{
			ID:     fmt.Sprintf("id%d", i),
			UserID: user.ID,
		}

		err := s.CreateNote(context.Background(), user, note)
		assert.NoError(t, err)
	}

	returnNotes, err := s.GetNotes(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(returnNotes))
}

func TestDeleteNote(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	noteID := "noteID"
	n := &models.Note{ID: noteID, UserID: user1.ID}

	s := NewNoteLocalStorage()

	err := s.CreateNote(context.Background(), user1, n)
	assert.NoError(t, err)

	err = s.DeleteNote(context.Background(), user1, noteID)
	assert.NoError(t, err)

	err = s.CreateNote(context.Background(), user1, n)
	assert.NoError(t, err)

	err = s.DeleteNote(context.Background(), user2, noteID)
	assert.Error(t, err)
	assert.Equal(t, err, note.ErrNoteNotFound)
}
