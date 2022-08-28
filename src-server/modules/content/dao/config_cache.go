// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/lache/driver"
	"go.uber.org/zap"
	"time"
)

func clearAllConfigCache() {
	boot.Lache.Delete("K-Content-Dao-Config-WebConfig")
}

func getWebConfig() (ContentWebConfigData, bool) {
	cnt := dao.SysConfigDao.CountByColumn("code", ContentWebConfigCode)
	if cnt == 0 {
		return ContentWebConfigData{}, false
	}
	mConfig := new(model.SysConfig)
	err := crud.DbSess().Where("code = ? and deleted = ?", ContentWebConfigCode, crud.FlagNo).First(mConfig).Error
	if err != nil {
		g3.ZL().Error("getWebConfig", zap.Error(err))
		return ContentWebConfigData{}, false
	}
	result := ContentWebConfigData{}
	err = jsoniter.UnmarshalFromString(mConfig.Value, &result)
	if err != nil {
		g3.ZL().Error("getWebConfig json", zap.Error(err))
		return result, false
	}
	return result, true
}

func getWebConfigFromCache() ContentWebConfigData {
	key := "K-Content-Dao-Config-WebConfig"
	result := ContentWebConfigData{}
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		_data, _ok := getWebConfig()
		if _ok {
			boot.Lache.Set(key, _data, driver.NotExpired)
		} else {
			boot.Lache.Set(key, _data, 10*time.Second)
		}
		return _data
	}
	return result
}
