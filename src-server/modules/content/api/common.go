package api

import "github.com/zhouhp1295/g3/net"

type updateMenuOrBannerParams struct {
	net.IdParams
	InMenu       string `json:"inMenu" form:"inMenu"`
	InMenuSort   int    `json:"inMenuSort" form:"inMenuSort"`
	InBanner     string `json:"inBanner" form:"inBanner"`
	InBannerSort int    `json:"inBannerSort" form:"inBannerSort"`
}
