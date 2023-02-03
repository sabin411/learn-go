package controllers

import (
	"go-gin/initializers"
	"go-gin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func Signup(ctx *gin.Context){
	var body struct{
		Email string `gorm:"unique"`
		Password string 
	}
	// 1. Get email and password from the body
	if ctx.BindJSON(&body) == nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	user := &models.User{Email: body.Email, Password: body.Password}
	// 2. hash the password. 
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Error hashing password"})
		return
	}
	// 3. Create the user. 
	result := initializers.DB.Create(&models.User{Email: user.Email, Password: string(hash)})
	if result.Error != nil {
		ctx.JSON(500, gin.H{"message": "Error creating user"})
		return
	}

	// 4. return response. 
	ctx.JSON(200, gin.H{"message": "User crated successfully"})
	return;
}


func Login(ctx *gin.Context){
	// 1. Get email and password from the body
	var loginBody struct {
		Email string 
		Password string
	}

	if ctx.BindJSON(&loginBody) == nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	// 2. Find the user by email
	user := models.User{}
	initializers.DB.First(&user, "email = ?", loginBody.Email)

	if user.ID == 0  {
		ctx.JSON(404, gin.H{"message": "Couldn't find the email"})
		return
	}
	// 3. Compare sent pass with saved hash pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))

	if err != nil {
		ctx.JSON(404, gin.H{"message": "Email or password is wrong."})
		return
	}else{
		
	}

	// 4. Generate JWT token


	// 5. return response
}