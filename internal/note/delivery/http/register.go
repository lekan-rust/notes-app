package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lekan-rust/notes-app/internal/note"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc note.UseCase) {
	h := NewHandler(uc)

	notes := router.Group("/notes")
	{
		notes.POST("", h.Create)
		notes.GET("", h.Get)
		notes.DELETE("", h.Delete)
	}
}
