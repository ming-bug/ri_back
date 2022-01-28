package utils

import (
	"github.com/dgrijalva/jwt-go"
	"ri/model"
	"time"
)

// 生成 token
func GenerateToken(hu model.HmacUser, autoLogin bool) (string, error) {
	// 定义过期时间，如果选择自动登录，token过期时间设置为一周，否则为一天
	expireTime := time.Time{}
	if autoLogin {
		expireTime = time.Now().Add(7 * 24 * time.Hour)
	}
	expireTime = time.Now().Add(24 * time.Hour)
	claims := &model.JwtClaims{
		UserUin:  hu.UserUin,
		Username: hu.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ri_server",
			Subject:   "auth_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedString, err := token.SignedString(model.JwtKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *model.JwtClaims, error) {
	claims := &model.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})
	return token, claims, err
}
