package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// Claim 定义token中的payload存储的信息
type Claim struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var secretKey = []byte("blkcor-secret")

// GenToken 生成jwt签名
func GenToken(userId int64, username string, email string) (string, error) {
	claims := Claim{
		userId,
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			Issuer:    "somebody",                                         // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析token信息
func ParseToken(tokenString string) (*Claim, error) {
	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "could not JSON decode header") {
			fmt.Println("token可能被篡改，头解析失败")
		}
		if err.Error() == "token signature is invalid: signature is invalid" {
			fmt.Println("令牌签名无效")
		}
		if err.Error() == "token has invalid claims: token is expired" {
			fmt.Println("令牌已过期")
		}
		return nil, err
	}

	if token.Valid {
		fmt.Println("令牌有效")
		return claims, nil
	}
	return nil, errors.New("无效的令牌")
}
