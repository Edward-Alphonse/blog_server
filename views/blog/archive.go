package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogArchiveHandler struct {
	context *gin.Context
}

func BlogArchive(context *gin.Context) {
	handler := NewBlogArchiveHandler(context)
	handler.Handle()
}

func NewBlogArchiveHandler(context *gin.Context) *BlogArchiveHandler {
	return &BlogArchiveHandler{
		context: context,
	}
}

func (h *BlogArchiveHandler) Handle() {
	h.makeResponse("")
}

func (h *BlogArchiveHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"tags": []string{"123", "456"},
		},
	})
}
