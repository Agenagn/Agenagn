package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createJWTToken(payload map[string]interface{}, secretKey string, tokenExpiration int64) (string, error) {
    // Create a new JWT token with the given payload and secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
    // Set the token expiration time to 1 hour from now
    token.Claims.(jwt.MapClaims)["exp"] = tokenExpiration
    // Sign the token with the secret key and return the signed token as a string
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func GetToken(userId int, role string ) (string, error){
	payload := map[string]interface{}{
		"sub": "12345",         // The user ID
		"iat": time.Now().Unix(),       // The token issue time (UNIX timestamp)
		"exp": time.Now().Add(time.Hour * 48).Unix(),  // The token expiration time (UNIX timestamp)
			
		"metadata": map[string]interface{}{
			"userId": userId,
			"role": role,
		},
	}
	tokenExpiration := time.Now().Add(time.Hour * 48).Unix()
	token, err := createJWTToken(payload, os.Getenv("JWT_SECRET"),tokenExpiration )
	if err != nil {
		return "", err
	}
	return token, nil
}
