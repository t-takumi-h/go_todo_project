package auth

import (
	"fmt"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := checkTokenFromHeader(r)
		if err != nil {
			fmt.Println("token err")
			return
		}
		next.ServeHTTP(w, r)
	})
}
