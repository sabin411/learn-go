package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string, secretKey []byte) (bool, error) {
    // Parse the JWT token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // Return the secret key
        return secretKey, nil
    })
    if err != nil {
        return false, err
    }
    // Check if the token is valid
    if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return true, nil
    } else {
        return false, fmt.Errorf("invalid token")
    }
}


