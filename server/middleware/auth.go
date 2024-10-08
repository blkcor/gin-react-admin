package middleware

import (
	"github.com/blkcor/gin-react-admin/enums"
	"github.com/blkcor/gin-react-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !IsUserLogin(context) {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    enums.ToLoginCode,
				"message": "Please login first!",
			})
			context.Abort()
			return
		}
		//user have login
		context.Next()
	}
}

func IsUserLogin(context *gin.Context) bool {
	//get jwt token from request header
	_, err := jwt.GetClaimFromContext(context)
	if err != nil {
		return false
	}
	return true
}
