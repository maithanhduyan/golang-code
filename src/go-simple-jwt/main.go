package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int
	Email    string
	Password string
}

var signingKey = []byte("your-secret-key")

func main() {
	user := User{
		ID:       1,
		Email:    "user@example.com",
		Password: "password123",
	}

	tokenString, err := generateJWT(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Token:", tokenString)

	// Verify and parse the JWT
	parsedUser, err := parseJWT(tokenString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", parsedUser.ID)
	fmt.Println("Email:", parsedUser.Email)
}

func generateJWT(user User) (string, error) {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseJWT(tokenString string) (*User, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &User{
			ID:       int(claims["id"].(float64)),
			Email:    claims["email"].(string),
			Password: "", // Exclude password from the parsed user object
		}
		return user, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
