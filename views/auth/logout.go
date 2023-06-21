package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(context *gin.Context) {
	fmt.Print("hello, world")
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}

type LogoutHandler struct {
	context *gin.Context
	// reqParams *LoginReqParams
}

func NewLogoutHandler(context *gin.Context) *LogoutHandler {
	return &LogoutHandler{
		context: context,
	}
}

func (h *LogoutHandler) Handle() {

}

func (h *LogoutHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": message,
	})
}
