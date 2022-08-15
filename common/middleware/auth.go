package middleware

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	cWrapper "github.com/fajarabdillahfn/merchants/common/wrapper"
)

var jwtKey = []byte("merchant")

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			log.Println("unauthorized, want Bearer token")
			cWrapper.ErrorJSON(w, errors.New("unauthorized, want Bearer token"), http.StatusUnauthorized)
			return
		}
		if !strings.Contains(authorizationHeader, "Bearer") {
			log.Println("unauthorized, want Bearer token")
			cWrapper.ErrorJSON(w, errors.New("unauthorized, want Bearer token"), http.StatusUnauthorized)
			return
		}
		jwtToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println("invalid token:", jwt.ErrSignatureInvalid)
				cWrapper.ErrorJSON(w, errors.New("invalid token"), http.StatusUnauthorized)
				return
			}
			log.Println("error token:", err.Error())
			cWrapper.ErrorJSON(w, errors.New("error token"), http.StatusBadRequest)
			return
		}
		if !token.Valid {
			log.Println("invalid token:", token)
			cWrapper.ErrorJSON(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		r.Header.Add("accessUserId", strconv.Itoa(claims.UserId))
		next.ServeHTTP(w, r)
	})
}
