package note

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, user *models.User, note *models.Note) error
	GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error)
	DeleteNote(ctx context.Context, user *models.User, id string) error
}
