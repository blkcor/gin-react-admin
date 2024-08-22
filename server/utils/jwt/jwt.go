package jwt

import (
	"errors"
	"fmt"
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// Claim 定义token中的payload存储的信息
type Claim struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint32 `json:"role_id"`
	RoleCode string `json:"role_code"`
	jwt.RegisteredClaims
}

// GenToken 生成jwt签名
func GenToken(userId uint32, username string, email string, roleCode string, roleId uint32, accessKey string) (string, error) {
	claims := Claim{
		userId,
		username,
		email,
		roleId,
		roleCode,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			Issuer:    "blkcor",                                           // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessKey))
}

// ParseToken 解析token信息
func ParseToken(tokenString string, accessKey string) (*Claim, error) {
	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessKey), nil
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
		return claims, nil
	}
	return nil, errors.New("无效的令牌")
}

// GetClaimFromContext 从gin.Context中获取token并解析
func GetClaimFromContext(ctx *gin.Context) (*Claim, error) {
	auth := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, errors.New("无效的令牌")
	}
	token := auth[7:]

	claim, err := ParseToken(token, section.AppConfig.AccessKey)
	if err != nil {
		return nil, err
	}
	return claim, nil
}
