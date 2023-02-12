package controllers

import (
	"go-gin/initializers"
	"go-gin/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)



func Signup(ctx *gin.Context){
	var body struct{
		Email string `gorm:"unique"`
		Password string 
	}
	// 1. Get email and password from the body
	if ctx.Bind(&body) != nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	// 2. hash the password. 
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Error hashing password"})
		return
	}
	// 3. Create the user. 
	result := initializers.DB.Create(&models.User{Email: body.Email, Password: string(hash)})
	if result.Error != nil {
		ctx.JSON(500, gin.H{"message": "Error creating user" })
		return
	}

	// 4. return response. 
	ctx.JSON(200, gin.H{"message": "User crated successfully"})

}



func CreateToken(expirationTime time.Time, jwtKey []byte, email string ) (string, error){
	type Claims struct {
		Email string `json:"username"`
		jwt.RegisteredClaims
	}

// 4. Generate JWT token
	
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}


func Login(ctx *gin.Context){
	// Create the JWT key used to create the signature
	accessKey := []byte(os.Getenv("ACCESS_SECRET_KEY"));
	refreshKey := []byte(os.Getenv("REFRESH_SECRET_KEY"));

	// 1. Get email and password from the body
	var loginBody struct {
		Email string 
		Password string
	}

	if ctx.Bind(&loginBody) != nil {
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
	}

	// 4. Generate JWT token
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	refreshTokenExpirationDuration := time.Now().Add(time.Hour * 24 * 30);
	accessTokenExpirationDuration := time.Now().Add(time.Hour * 24);

	refreshToken, refreshErr := CreateToken(refreshTokenExpirationDuration, refreshKey, user.Email)
	if(refreshErr != nil){
		ctx.JSON(500, gin.H{"message": "Error creating refresh token"})
		return
	}
	accessToken, accessErr  := CreateToken(accessTokenExpirationDuration, accessKey, user.Email)
	if(accessErr != nil){
		ctx.JSON(500, gin.H{"message": "Error creating access token"})
		return
	}


	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	ctx.SetCookie("refresh_token", refreshToken, int(refreshTokenExpirationDuration.Unix()), "/", "localhost", false, true)
	ctx.SetCookie("access_token", accessToken, int(accessTokenExpirationDuration.Unix()), "/", "localhost", false, true)

	ctx.JSON(200, gin.H{"message": "User logged in successfully"})
}

func Welcome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Welcome to the home page"})
	
}