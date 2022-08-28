// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package utils

import (
	"os"
	"runtime"
)

// IsFile returns true if given path exists as a file (i.e. not a directory).
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir returns true if given path is a directory, and returns false when it's
// a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// IsExist returns true if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsMac() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

func IsWin() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}
