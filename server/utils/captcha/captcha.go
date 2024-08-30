package captcha

import (
	"github.com/afocus/captcha"
	"github.com/blkcor/gin-react-admin/core/cache"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/gin-gonic/gin"
)

// GenerateCaptchaAndImg 生成验证码和对应的图片
func GenerateCaptchaAndImg(context *gin.Context) (*captcha.Image, string) {
	c := captcha.New()
	// 设置字体
	err := c.SetFont("fonts/comic.ttf")
	if err != nil {
		logger.Error("Captcha font error: ", err)
		context.JSON(500, response.BaseResponse[any]{
			Success: false,
			Message: "Captcha font error",
			Data:    nil,
		})
	}
	return c.Create(4, captcha.NUM)
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(context *gin.Context, captchaCode string) bool {
	clientIP, ok := context.Get("ClientIP")
	if !ok {
		logger.Error("Get client ip failed")
		return false
	}
	cpt, err := cache.RDB.Get(context, clientIP.(string)).Result()
	if err != nil {
		logger.Error("Get captcha code from redis failed: ", err)
		return false
	}
	if cpt != captchaCode {
		logger.Error("Captcha verification failed")
		return false
	}
	return true
}
