package service

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3/helpers"
	"io/ioutil"
	"mime/multipart"
	"path"
	"runtime"
)

type uploadService struct {
}

var UploadService = new(uploadService)

func (service *uploadService) UploadImage(file *multipart.FileHeader) (filePath string, msg string, ok bool) {
	fileType := file.Header.Get("Content-Type")
	if fileType != "image/png" && fileType != "image/jpeg" && fileType != "image/x-icon" {
		msg = "文件格式错误, 仅支持: image/png、image/jpeg、image/x-icon"
		return
	}

	fileSize := float32(file.Size / 1024 / 1024)
	if fileSize > boot.StorageCfg.MaxSize {
		msg = fmt.Sprintf("上传图片大小不能超过 %fMB!", boot.StorageCfg.MaxSize)
		return
	}

	saveFilename, err := SaveUploadFile(file)
	if err != nil {
		msg = err.Error()
		return
	}
	filePath = boot.StorageCfg.UriPrefix + saveFilename
	ok = true
	return
}

var ImagesTypes = []string{
	"image/png",
	"image/jpeg",
	"image/x-icon",
}

var DefaultFileSuffix = map[string]string{
	"image/png":    ".png",
	"image/jpeg":   ".jpg",
	"image/x-icon": ".ico",
}

func SaveUploadFile(file *multipart.FileHeader) (string, error) {
	fileType := file.Header.Get("Content-Type")
	categoryPath := ""
	filename := ""
	suffix := path.Ext(file.Filename)

	if helpers.IndexOf[string](ImagesTypes, fileType) != -1 {
		categoryPath = "images"
	} else {
		categoryPath = "files"
	}
	if len(suffix) == 0 {
		if _suffix, ok := DefaultFileSuffix[fileType]; ok {
			suffix = _suffix
		}
	}
	if len(suffix) == 0 {
		return filename, errors.New("未定义的文件后缀, 格式:" + fileType)
	}
	//
	src, err := file.Open()
	if err != nil {
		return filename, err
	}
	defer func() {
		_ = src.Close()
		runtime.GC()
	}()

	body, err := ioutil.ReadAll(src)
	if err != nil {
		return filename, err
	}
	_, _ = src.Seek(0, 0)

	md5Str := fmt.Sprintf("%x", md5.Sum(body))

	md5Dir := md5Str[:2]

	name := md5Str + suffix

	filename = path.Join("upload", categoryPath, md5Dir, name)

	_, err = boot.Storager.Write(filename, src, file.Size)
	if err != nil {
		return filename, err
	}
	return filename, nil
}
