package model

import "github.com/dgrijalva/jwt-go"

// 签名包含的参数
type JwtClaims struct {
	UserUin  string `json:"userUin"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type HmacUser struct {
	UserUin  string `json:"userUin"`
	Username string `json:"username"`
}

var JwtKey = []byte("ri")
