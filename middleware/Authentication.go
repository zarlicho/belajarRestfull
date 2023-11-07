package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func Auth(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tokenCookie, err := r.Cookie("MyCookie")
		if err != nil {
			if err == http.ErrNoCookie {
				// Cookie tidak ditemukan
				fmt.Println("Cookie tidak ditemukan.")
			} else {
				// Terjadi kesalahan lain
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		// Anda dapat mengakses nilai cookie dengan cookie.Value
		tokenString := tokenCookie.Value
		fmt.Printf("Nilai Cookie: %s\n", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("asc22332afg0061729940qqr"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return
			}
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}
