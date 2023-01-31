package main

import (
	"go-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default();

	r.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Hello Worlds",
		})
	})

	r.Run()
}