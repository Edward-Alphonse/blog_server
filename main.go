package main

import (
	"fmt"

	"blog_server/consts"
	"blog_server/views/auth"
	"blog_server/views/blog"
	"blog_server/views/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print(consts.RegisterStatusMap)
	router := gin.Default()
	router.POST("/blog/list", blog.BlogList)
	router.GET("/blog/tags", blog.BlogTags)
	router.GET("/blog/detail", blog.BlogDetail)
	router.GET("/blog/collection", blog.BlogCollection)
	router.GET("/blog/archive", blog.BlogArchive)
	router.GET("/profile/main", profile.ProfileMain)
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/logout", auth.Logout)
	router.POST("/auth/register", auth.Register)
	router.Run(":5000")
}
