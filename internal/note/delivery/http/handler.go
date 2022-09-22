package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lekan-rust/notes-app/internal/auth"
	"github.com/lekan-rust/notes-app/internal/models"
	"github.com/lekan-rust/notes-app/internal/note"
)

type Note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Handler struct {
	useCase note.UseCase
}

func NewHandler(useCase note.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateNote(c.Request.Context(), user, inp.Title, inp.Description); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Notes []*Note `json:"notes"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	nts, err := h.useCase.GetNotes(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Notes: toNotes(nts),
	})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(gin.AuthUserKey).(*models.User)

	if err := h.useCase.DeleteNote(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toNotes(ns []*models.Note) []*Note {
	out := make([]*Note, len(ns))

	for i, n := range ns {
		out[i] = toNote(n)
	}

	return out
}

func toNote(n *models.Note) *Note {
	return &Note{
		ID:          n.ID,
		Title:       n.Title,
		Description: n.Description,
	}
}
