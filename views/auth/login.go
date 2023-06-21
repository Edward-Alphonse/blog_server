package auth

import (
	"fmt"
	"net/http"

	"blog_server/consts"
	"blog_server/db"
	"blog_server/models"
	"blog_server/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type LoginHandler struct {
	context *gin.Context
	// reqParams *LoginReqParams
}

func Login(context *gin.Context) {
	handler := NewLoginHandler(context)
	handler.Handle()
}

func NewLoginHandler(context *gin.Context) *LoginHandler {
	return &LoginHandler{
		context: context,
	}
}

func (h *LoginHandler) Handle() {
	params, err := utils.ExtraReqParams(h.context)
	if err != nil {
		message := fmt.Sprintf("err: %v", err)
		h.makeResponse(message)
		return
	}
	if utils.IsEmptyEmail(params.Email) {
		message := consts.LoginStatusMap[1]
		h.makeResponse(message)
		return
	}
	if utils.IsPassword(params.Password) {
		message := consts.LoginStatusMap[2]
		h.makeResponse(message)
		return
	}

	sql := queryUserSql(params)
	user, err := db.Instance.Query(sql)
	if err != nil {
		fmt.Println("query err: ", err)
		message := consts.LoginStatusMap[5]
		h.makeResponse(message)
		return
	}
	if user.Email != params.Email || user.Password != params.Password {
		message := consts.LoginStatusMap[4]
		h.makeResponse(message)
		return
	}
	message := consts.LoginStatusMap[0]
	h.makeResponse(message)
}

func (h *LoginHandler) makeResponse(message string) {
	cookie := &http.Cookie{
		Name:     "__cookie1",
		Value:    "hzc",
		HttpOnly: true,
	}
	http.SetCookie(h.context.Writer, cookie)

	h.context.SetCookie("__cookie2", "hzc", 1, "path", "", false, true)
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": message,
	})
}

func queryUserSql(params *models.AuthParams) string {
	return fmt.Sprintf("select * from users where email='%s' limit 1", params.Email)
}
