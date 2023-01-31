package main

import (
	"go-gin/controllers"
	"go-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default();

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsFindAll)
	r.GET("/post/:id", controllers.PostFindOne)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run()
}