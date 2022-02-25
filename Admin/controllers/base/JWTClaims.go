package base

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID   int `json:"user_id"`
	UserName string
}
