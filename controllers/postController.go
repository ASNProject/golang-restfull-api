package controllers

import (
	"asnproject/backend-api/models"

	"github.com/gin-gonic/gin"
)

// get all post
func FindPosts(c *gin.Context) {
	// get data from database using model
	var posts []models.Post
	models.DB.Find(&posts)

	// return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "List Data Posts",
		"data": posts,
	})
}