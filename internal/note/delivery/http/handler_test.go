package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lekan-rust/notes-app/internal/auth"
	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/lekan-rust/notes-app/internal/note/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()

	group := r.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.NoteUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &createInput{
		Title:       "testtitle",
		Description: "testdescription",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreateNote", testUser, inp.Title, inp.Description).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/notes", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()

	group := r.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.NoteUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	nts := make([]*models.Note, 5)
	for i := 0; i < 5; i++ {
		nts[i] = &models.Note{
			ID:          "id",
			Title:       "title",
			Description: "description",
		}
	}

	uc.On("GetNotes", testUser).Return(nts, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notes", nil)
	r.ServeHTTP(w, req)

	expectedOut := &getResponse{Notes: toNotes(nts)}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestDelete(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.NoteUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &deleteInput{
		ID: "id",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("DeleteNote", testUser, inp.ID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/notes", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
