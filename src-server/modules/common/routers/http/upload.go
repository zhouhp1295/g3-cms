// Copyright (c) 554949297@qq.com . 2022-2022. All rights reserved

//go:build http
// +build http

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/service"
	"github.com/zhouhp1295/g3/net"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/common/uploadImage", handleUploadImage)
	})
}

func handleUploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		g3.ZL().Error("can't load file", zap.Error(err))
		net.FailedMessage(ctx, err.Error())
		return
	}
	filePath, msg, ok := service.UploadService.UploadImage(file)
	if !ok {
		g3.ZL().Error("can't upload file", zap.String("msg", msg), zap.Error(err))
		net.FailedMessage(ctx, msg)
		return
	}
	net.SuccessData(ctx, gin.H{"url": filePath})
}
