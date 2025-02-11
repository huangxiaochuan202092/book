package main

import (
	"book-management-api/database"
	"book-management-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	// 修改静态资源路径
	r.Static("/static/", "./static/")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// 路由
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook) // 添加删除图书路由

	r.Run(":8080")
}