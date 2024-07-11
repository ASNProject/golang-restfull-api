// Program utama ditandai dengan package main
package main

// Import library
import (
	"github.com/gin-gonic/gin"
)

// Membuat fungsi
func main() {
	//inisialisasi Gin
	router := gin.Default()

	//Membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//mulai server dengan port 3000
	router.Run(":3000")
}