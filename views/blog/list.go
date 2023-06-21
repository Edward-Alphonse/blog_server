package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogListHandler struct {
	context *gin.Context
}

func BlogList(context *gin.Context) {
	handler := NewBlogListHandler(context)
	handler.Handle()
}

func NewBlogListHandler(context *gin.Context) *BlogListHandler {
	return &BlogListHandler{
		context: context,
	}
}

func (h *BlogListHandler) Handle() {
	h.makeResponse("")
}

func (h *BlogListHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
