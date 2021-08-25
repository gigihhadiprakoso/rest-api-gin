package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"strings"
	. "rest-api-gin/models"
)

type Claims struct{
	Username string `json:"username"`
	UserID uint `json:"user_id"`
	CompanyID uint `json:"company_id"`
	jwt.StandardClaims
}


func CreateToken(user Users) (string, error) {
	claims := &Claims{
		Username: user.Username,
		UserID: user.ID,
		CompanyID: user.CompanyID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("gigih-rupawan"))
}

func ExtractToken(auth string) string {
	arrToken := strings.Split(auth, " ")

	if len(arrToken) ==2{
		return arrToken[1]
	}
	return ""
}