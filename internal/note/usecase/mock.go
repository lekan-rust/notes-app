package usecase

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/stretchr/testify/mock"
)

type NoteUseCaseMock struct {
	mock.Mock
}

func (m NoteUseCaseMock) CreateNote(ctx context.Context, user *models.User, title, description string) error {
	args := m.Called(user, title, description)

	return args.Error(0)
}

func (m NoteUseCaseMock) GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error) {
	args := m.Called(user)

	return args.Get(0).([]*models.Note), args.Error(1)
}

func (m NoteUseCaseMock) DeleteNote(ctx context.Context, user *models.User, id string) error {
	args := m.Called(user, id)

	return args.Error(0)
}
