package utils

import (
	"blog_server/models"

	"github.com/gin-gonic/gin"
)

func ExtraReqParams(context *gin.Context) (*models.AuthParams, error) {
	reqParams := models.AuthParams{}
	if err := context.BindJSON(&reqParams); err != nil {
		return nil, err
	}
	return &reqParams, nil
}

func IsEmptyEmail(email string) bool {
	return len(email) == 0
}

func IsPassword(password string) bool {
	return len(password) == 0
}
