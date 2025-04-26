package jwtutil

import (
	"errors"
	"scaffold-demo/config"
	"scaffold-demo/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

// JwtCustomClaims 自定义Claims类型，用于存储用户信息和JWT的标准声明
type JwtCustomClaims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 封装生成token的方法
func GenerateToken(username string) (string, error) {
	claims := JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JwtExpireTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bobbybai",
			Subject:   "bobbybai",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// ParseToken 封装解析token的方法
func ParseToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		// 解析失败，返回错误信息
		logs.Error(nil, "解析token失败")
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil	
	} else {
		// 解析失败，返回错误信息
		logs.Error(nil, "token不合法")
		return nil, errors.New("token不合法")	
	}

}
