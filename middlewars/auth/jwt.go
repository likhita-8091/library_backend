package auth

import (
	"github.com/golang-jwt/jwt"
)

var Secret = "jw"

// 定义一个自定义jwt claims
type MyJwtClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
	jwt.StandardClaims
}

// 签发一个token
func (j MyJwtClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	return token.SignedString([]byte(Secret))
}
