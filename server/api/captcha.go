package api

import (
	"github.com/blkcor/gin-react-admin/core/cache"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/utils/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"time"
)

// Captcha 验证码
func Captcha(context *gin.Context) {
	img, str := captcha.GenerateCaptchaAndImg(context)
	// store the captcha code in the redis
	clientIP, ok := context.Get("ClientIP")
	if !ok {
		context.JSON(500, gin.H{
			"message": "Get client ip failed",
		})
	}
	//设置过期时间为一分钟
	cache.RDB.Set(context, clientIP.(string), str, time.Minute)
	logger.Info("获取到验证码为:", str)
	//设置响应头，将验证码图片响应出去
	context.Writer.Header().Set("Content-Type", "image/png")
	_ = png.Encode(context.Writer, img)
}
