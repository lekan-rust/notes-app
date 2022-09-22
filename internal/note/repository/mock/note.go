package mock

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/stretchr/testify/mock"
)

type NoteStorageMock struct {
	mock.Mock
}

func (s *NoteStorageMock) CreateNote(ctx context.Context, user *models.User, nt *models.Note) error {
	args := s.Called(user, nt)

	return args.Error(0)
}

func (s *NoteStorageMock) GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error) {
	args := s.Called(user)
	return args.Get(0).([]*models.Note), args.Error(1)
}

func (s *NoteStorageMock) DeleteNote(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}
