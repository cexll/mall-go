package jwtx

import (
	"fmt"
	"mall-go/common/conf"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	Secret string
}

func NewJwt(c conf.JwtAuth) *Jwt {
	return &Jwt{
		Secret: c.Secret,
	}
}

func (t *Jwt) GenerateJwt(uid int64) (string, error) {
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "mall-go",                                             // 签发人
		"iat": now,                                                   // 签发时间
		"exp": now + int64(7200),                                     // 过期时间
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // 什么时间之前不可用
		"uid": strconv.FormatInt(uid, 10),
	})
	tokenString, err := token.SignedString([]byte(t.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

var hmacSampleSecret []byte

func (t *Jwt) CheckJwt(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
