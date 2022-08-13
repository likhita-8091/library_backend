package auth

import "github.com/golang-jwt/jwt"

var Secret = "jw"

// 定义一个自定义jwt claims
type MyJwtClaims struct {
	Name  string `json:"name"`
	Admin string `json:"admin"`
	jwt.StandardClaims
}

// 签发一个token
func (j MyJwtClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	return token.SignedString([]byte(Secret))
}

// 解析token
func (j MyJwtClaims) ParseToken() {

}
