package main

import (
	"github.com/gin-gonic/gin"
	"go-openapi3-sample/routes"
)

func main() {
	blog := &routes.Blog{}

	router := gin.Default()

	router.GET("/blogs", blog.GetBlogs)
	router.GET("/blogs/:id", blog.GetBlog)
	router.POST("/blogs", blog.CreateBlog)
	router.PUT("/blogs/:id", blog.UpdateBlog)
	router.DELETE("/blogs/:id", blog.DeleteBlog)

	router.Run(":8080")
}
