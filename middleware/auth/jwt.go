package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

const (
	secretKey = "testSecretKey"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "test"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(secretKey))

	w.Write([]byte(tokenString))
})

func checkTokenFromHeader(r *http.Request) (*jwt.Token, error){
	auth := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(auth, "Bearer ")
	fmt.Println("test1",tokenString)
	verifyToken, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	return verifyToken, nil
}

func verifyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
