package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogCollectionHandler struct {
	context *gin.Context
}

func BlogCollection(context *gin.Context) {
	handler := NewBlogCollectionHandler(context)
	handler.Handle()
}

func NewBlogCollectionHandler(context *gin.Context) *BlogCollectionHandler {
	return &BlogCollectionHandler{
		context: context,
	}
}

func (h *BlogCollectionHandler) Handle() {
	h.makeResponse("")
}

func (h *BlogCollectionHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"tags": []string{"123", "456"},
		},
	})
}
