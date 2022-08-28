// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package utils

func YesOrNo(s string) bool {
	if s == "0" {
		return false
	}
	return true
}

func TFStrVal(cond bool, t, f string) string {
	if cond {
		return t
	}
	return f
}

func DefaultStrVal(v, v2 string) string {
	if len(v) > 0 {
		return v
	}
	return v2
}

func TFIntVal(cond bool, t, f int) int {
	if cond {
		return t
	}
	return f
}
