package middleware

import (
	"time"

	"go_Rahman_Wahabi_Hasibuan/22_Middleware/constants"

	"github.com/dgrijalva/jwt-go"
)

func CreateBookToken(bookId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["bookId"] = bookId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}
