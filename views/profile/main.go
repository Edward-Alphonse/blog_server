package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileMainHandler struct {
	context *gin.Context
}

func ProfileMain(context *gin.Context) {
	handler := NewProfileMainHandler(context)
	handler.Handle()
}

func NewProfileMainHandler(context *gin.Context) *ProfileMainHandler {
	return &ProfileMainHandler{
		context: context,
	}
}

func (h *ProfileMainHandler) Handle() {

}

func (h *ProfileMainHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": message,
	})
}
