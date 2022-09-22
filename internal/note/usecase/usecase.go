package usecase

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/lekan-rust/notes-app/internal/note"
)

type NoteUseCase struct {
	noteRepo note.NoteRepository
}

func NewNoteUseCase(noteRepo note.NoteRepository) *NoteUseCase {
	return &NoteUseCase{
		noteRepo: noteRepo,
	}
}

func (n NoteUseCase) CreateNote(ctx context.Context, user *models.User, title, description string) error {
	note := &models.Note{
		Title:       title,
		Description: description,
	}

	return n.noteRepo.CreateNote(ctx, user, note)
}

func (n NoteUseCase) GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error) {
	return n.noteRepo.GetNotes(ctx, user)
}

func (n NoteUseCase) DeleteNote(ctx context.Context, user *models.User, id string) error {
	return n.noteRepo.DeleteNote(ctx, user, id)
}
