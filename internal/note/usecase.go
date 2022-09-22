package note

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
)

type UseCase interface {
	CreateNote(ctx context.Context, user *models.User, title, description string) error
	GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error)
	DeleteNote(ctx context.Context, user *models.User, id string) error
}
