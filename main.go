package main

import (
	"go-gin/bootstrap"
	"go-gin/initializers"

	"go.uber.org/fx"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	// r := gin.Default();

	// r.POST("/sign-up", controllers.Signup)
	// r.POST("/login", controllers.Login);
	// r.GET("/home", middleware.AuthMiddleware ,controllers.Welcome);
	// r.Run()
	fx.New(bootstrap.Module).Run()

}