package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Unauthorized(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "其实走不到这里",
	})
}
