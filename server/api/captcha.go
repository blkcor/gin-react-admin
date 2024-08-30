package api

import (
	"github.com/blkcor/gin-react-admin/core/cache"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/utils/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"time"
)

// Captcha godoc
//
//	@Summary		获取验证码
//	@Description	获取验证码接口
//	@Tags			登录相关接口
//	@Produce		image/png
//	@Success		200	{string}	string	"ok"
//	@Failure		500	{string}	string	"error"
//	@Router			/captcha [get]
func Captcha(context *gin.Context) {
	img, str := captcha.GenerateCaptchaAndImg(context)
	// store the captcha code in the redis
	clientIP, ok := context.Get("ClientIP")
	if !ok {
		context.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "Get client ip failed",
			Data:    nil,
		})
	}
	//设置过期时间为一分钟
	cache.RDB.Set(context, clientIP.(string), str, time.Minute)
	logger.Info("客户端IP:", clientIP)
	logger.Info("获取到验证码为:", str)
	//设置响应头，将验证码图片响应出去
	context.Writer.Header().Set("Content-Type", "image/png")
	_ = png.Encode(context.Writer, img)
}
