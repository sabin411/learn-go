package controllers

import (
	"go-gin/initializers"
	"go-gin/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context){
	var body struct{
		Title string `json:"title" binding:"required"`
		Body string `json:"body" binding:"required"`
	}
	ctx.Bind(&body)
	// Get data off request body create a post and return it. 
	post := models.POST{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if(result.Error != nil){
		ctx.JSON(500, gin.H{"message": "Error creating post"})
		return
	} 
	ctx.JSON(200, gin.H{"message": "Post created successfully", "post": post})
	return
	
}

func PostsFindAll(ctx *gin.Context){
	var posts []models.POST
	initializers.DB.Find(&posts)
	ctx.JSON(200, gin.H{"message": "All posts", "posts": posts})
	return
}

func PostFindOne(ctx *gin.Context){
	var post models.POST
	id := ctx.Param("id")

	initializers.DB.First(&post, id)
	ctx.JSON(200, gin.H{"message": "Post found", "post": post})
	return;
}

func PostUpdate(ctx *gin.Context){
	var post models.POST
	id := ctx.Param("id")

	var body struct{
		Title string `json:"title" binding:"required"`
		Body string `json:"body" binding:"required"`
	}
	ctx.Bind(&body)

	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(&models.POST{Title: body.Title, Body: body.Body})
	ctx.JSON(200, gin.H{"message": "Post updated", "post": post})
	return;
}

func PostDelete(ctx *gin.Context){
	id := ctx.Param("id")

	if(initializers.DB.First(&models.POST{}, id).RowsAffected == 0){
		ctx.JSON(404, gin.H{"message": "Post not found"})
		return
	}

	initializers.DB.Delete(&models.POST{}, id)

	ctx.JSON(200, gin.H{"message": "Post deleted successfully"})
	return;
}