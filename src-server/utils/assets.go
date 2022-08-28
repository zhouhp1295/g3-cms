// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3/helpers"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"runtime"
)

var ImagesTypes = []string{
	"image/png",
	"image/jpeg",
}

var DefaultFileSuffix = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
}

func SaveUploadFile(uploadRoot string, file *multipart.FileHeader) (string, error) {
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

	md5Str := fmt.Sprintf("%x", md5.Sum(body))

	md5Dir := md5Str[:2]

	dir := path.Join(uploadRoot, "upload", categoryPath, md5Dir)

	if !IsExist(dir) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return filename, err
		}
	}
	name := md5Str + suffix
	dst := path.Join(dir, name)
	filename = path.Join("upload", categoryPath, md5Dir, name)

	if IsExist(dst) {
		return filename, nil
	}

	err = ioutil.WriteFile(dst, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filename, nil
}
