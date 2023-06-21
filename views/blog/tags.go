package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogTagsHandler struct {
	context *gin.Context
}

func BlogTags(context *gin.Context) {
	handler := NewBlogTagsHandler(context)
	handler.Handle()
}

func NewBlogTagsHandler(context *gin.Context) *BlogTagsHandler {
	return &BlogTagsHandler{
		context: context,
	}
}

func (h *BlogTagsHandler) Handle() {
	h.makeResponse("")
}

func (h *BlogTagsHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"tags": []string{"123", "456"},
		},
	})
}
