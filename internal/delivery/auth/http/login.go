package http

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	cWrapper "github.com/fajarabdillahfn/merchants/common/wrapper"
)

var jwtKey = []byte("merchant")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Response struct {
	Token string `json:"access_token"`
}

func (d *Delivery) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := d.authUC.UserGetByUsername(ctx, r.FormValue("username"))
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	if hashedPassword != user.Password {
		cWrapper.ErrorJSON(w, err)
		return
	}

	expirationTime := time.Now().Add(25 * time.Minute)
	claims := &Claims{
		UserId:   user.Id,
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, Response{Token: tokenString}, "token")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}
