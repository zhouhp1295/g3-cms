package dao

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/lache/driver"
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
		boot.Logger.Warn("getWebConfig db err %s", err.Error())
		return ContentWebConfigData{}, false
	}
	result := ContentWebConfigData{}
	err = jsoniter.UnmarshalFromString(mConfig.Value, &result)
	if err != nil {
		boot.Logger.Warn("getWebConfig json err %s", err.Error())
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
