package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go-openapi3-sample/routes"
)

func main() {
	blog := &routes.Blog{}

	router := gin.Default()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})
	router.Use(corsConfig)

	router.GET("/blogs", blog.GetBlogs)
	router.GET("/blogs/:id", blog.GetBlog)
	router.POST("/blogs", blog.CreateBlog)
	router.PUT("/blogs/:id", blog.UpdateBlog)
	router.DELETE("/blogs/:id", blog.DeleteBlog)

	router.Run(":8080")
}
