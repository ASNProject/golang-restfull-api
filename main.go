package main

import (
	"asnproject/backend-api/controllers"
	"asnproject/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialisasi Gin
	router := gin.Default()

	// panggil koneksi database
	models.ConnectDatabase()

	//Membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// membuat route get all posts
	router.GET("/api/post", controllers.FindPosts)

	// membuat route store post
	router.POST("/api/post", controllers.StorePost)

	//mulai server dengan port 3000
	router.Run(":3000")
}