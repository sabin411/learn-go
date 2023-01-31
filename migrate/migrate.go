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
	// Migrate the schema
  initializers.DB.AutoMigrate(&models.POST{Title: "Blog 1"})
}
