package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/service"
)

func init() {
	boot.ApiPost("/common/uploadImage", handleUploadImage)
}

func handleUploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		boot.Logger.Error("handleUploadImage err = %s\n", err.Error())
		CommonApi.FailedMessage(ctx, err.Error())
		return
	}
	filePath, msg, ok := service.UploadService.UploadImage(file)
	if !ok {
		boot.Logger.Error("handleUploadImage err = %s\n", msg)
		CommonApi.FailedMessage(ctx, msg)
		return
	}
	CommonApi.SuccessData(ctx, gin.H{
		"url": filePath,
	})
}
