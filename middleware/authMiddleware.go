package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context){
	// get the bearer authorization from the header. 
	// if the header is not present, return 401.

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	// if the token is valid, check if they are a valid token 

	// isValid, err := utils.VerifyToken(accessToken, []byte(os.Getenv("ACCESS_SECRET_KEY")))
	// if(err != nil){
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"message": err})
	// 	return
	// }


	// if the token is invalid, return 401.
	fmt.Println("Auth middleware called", refreshToken, accessToken)

	ctx.Next()
}