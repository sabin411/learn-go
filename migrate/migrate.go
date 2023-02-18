package main

import (
	"go-gin/initializers"
	"go-gin/models"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
  initializers.DB.AutoMigrate(&models.User{})
}
