package auth

import (
	"blog_server/consts"
	"blog_server/db"
	"blog_server/models"
	"blog_server/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	handler := NewRegisterHandler(context)
	handler.Handle()
}

type RegisterHandler struct {
	context *gin.Context
}

func NewRegisterHandler(ctx *gin.Context) *RegisterHandler {
	return &RegisterHandler{
		context: ctx,
	}
}

func (h *RegisterHandler) Handle() {
	params, err := utils.ExtraReqParams(h.context)
	if err != nil {
		fmt.Println("err: ", err)
		h.makeResponse(consts.RegisterStatusMap[3])
		return
	}
	if utils.IsEmptyEmail(params.Email) || utils.IsPassword(params.Password) {
		h.makeResponse(consts.RegisterStatusMap[1])
		return
	}

	exist, err := h.isAccountExist(params)
	if err != nil {
		h.makeResponse(consts.RegisterStatusMap[3])
		return
	}
	if exist {
		h.makeResponse(consts.RegisterStatusMap[2])
		return
	}
	if err := h.register(params); err != nil {
		h.makeResponse(consts.RegisterStatusMap[3])
		return
	}
	h.makeResponse(consts.RegisterStatusMap[0])
}

func (h *RegisterHandler) isAccountExist(params *models.AuthParams) (bool, error) {
	sql := queryCountSql(params)
	count, err := db.Instance.QueryCount(sql)
	if err != nil {
		fmt.Println("err: ", err)
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

func (h *RegisterHandler) register(params *models.AuthParams) error {
	user := models.User{
		Id:       utils.GenerateId(),
		Name:     "0000000",
		Email:    params.Email,
		Password: params.Password,
	}

	sql := fmt.Sprintf("INSERT INTO %s values(0, %v, '%s', '%s', '%s')", "users", user.Id, user.Name, user.Password, user.Email)
	if err := db.Instance.Insert(sql); err != nil {
		return err
	}
	return nil
}

func (h *RegisterHandler) makeResponse(message string) {
	h.context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": message,
	})
}

func queryCountSql(params *models.AuthParams) string {
	return fmt.Sprintf("select count(email) from users where email = '%s'", params.Email)
}
