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
	
	r.POST("/sign-up", controllers.Signup)
	r.POST("/login", controllers.Login);
	r.Run()
}