package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogDetailHandler struct {
	context *gin.Context
}

func BlogDetail(context *gin.Context) {
	handler := NewBlogDetailHandler(context)
	handler.Handle()
}

func NewBlogDetailHandler(context *gin.Context) *BlogDetailHandler {
	return &BlogDetailHandler{
		context: context,
	}
}

func (h *BlogDetailHandler) Handle() {
	h.makeResponse("")
}

func (h *BlogDetailHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"tags": []string{"123", "456"},
		},
	})
}
