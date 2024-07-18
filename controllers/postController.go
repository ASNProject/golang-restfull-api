package controllers

import (
	"asnproject/backend-api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type validation post input
type ValidatePostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

// function get error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

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

// store a post
func StorePost(c *gin.Context) {
	// validate input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
						out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			}
			return
	}
	// Create post
	post := models.Post{
		Title: input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	//return response json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Create Successfully",
		"data": post,
	})
}

func FindPostsById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By ID : " + c.Param("id"),
		"data": post,
	})
}

// update post
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	// update post
	models.DB.Model(&post).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Update Successfully",
		"data": post,
	})
}