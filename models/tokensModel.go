package models

import(
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type Tokens struct{
	Hash string
	Token string
	IsDeleted int8
}


func CreateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("gigih-rupawan"))
}