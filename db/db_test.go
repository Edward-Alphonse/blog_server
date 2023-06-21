package db

import (
	"blog_server/models"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := NewDefaultDB()
	db.setupConnection()
	// db.Insert(models.User{
	// 	Id:       1,
	// 	Name:     "hzc",
	// 	Password: "12345",
	// 	Email:    "qq.com",
	// })
	user := &models.User{}
	if err := db.QueryOne(user); err != nil {
		fmt.Println("query one err: ", err)
	}
	fmt.Println(user)
}
