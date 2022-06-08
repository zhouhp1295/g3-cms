package dao

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zhouhp1295/g3-cms/boot"
	dao2 "github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
)

const ContentWebConfigCode = "content.web.config"

type contentConfigDAO struct {
	crud.BaseDao
}

var ContentConfigDao = &contentConfigDAO{
	crud.BaseDao{Model: new(model.SysConfig)},
}

type ContentWebConfigData struct {
	Title                 string `json:"title" form:"title"`                                 //网站标题
	Host                  string `json:"host" form:"host"`                                   //网站host
	Keywords              string `json:"keywords" form:"keywords"`                           //网站关键字
	Description           string `json:"description" form:"description"`                     //网站描述
	Copyright             string `json:"copyright" form:"copyright"`                         //网站底部Copyright
	Favicon               string `json:"favicon" form:"favicon"`                             //favicon
	Logo                  string `json:"logo" form:"logo"`                                   //logo
	Beian                 string `json:"beian" form:"beian"`                                 //备案号
	GonganBeian           string `json:"gonganBeian" form:"gonganBeian"`                     //公安网备案号
	ArticleSuffix         string `json:"articleSuffix" form:"articleSuffix"`                 //文章末尾统一追加内容
	Robots                string `json:"robots" form:"robots"`                               //网站robots文件
	BaiduSiteVerification string `json:"baiduSiteVerification" form:"baiduSiteVerification"` //百度站点验证
}

func (dao *contentConfigDAO) WebConfig() ContentWebConfigData {
	return getWebConfigFromCache()
}

func (dao *contentConfigDAO) UpdateWebConfig(data ContentWebConfigData, operator int64) (msg string, ok bool) {
	cnt := dao2.SysConfigDao.CountByColumn("code", ContentWebConfigCode)
	mConfig := new(model.SysConfig)
	if cnt > 0 {
		err := crud.DbSess().Where("code = ? and deleted = ?", ContentWebConfigCode, crud.FlagNo).First(mConfig).Error
		if err != nil {
			boot.Logger.Warn("UpdateWebConfig db err %s", err.Error())
			return err.Error(), false
		}
	} else {
		mConfig.Code = ContentWebConfigCode
	}
	var err error
	mConfig.Value, err = jsoniter.MarshalToString(data)
	if err != nil {
		boot.Logger.Warn("UpdateWebConfig MarshalToString err %s", err.Error())
		return err.Error(), false
	}
	if cnt > 0 {
		err = crud.DbSess().Updates(mConfig).Error
	} else {
		err = crud.DbSess().Create(mConfig).Error
	}
	if err != nil {
		boot.Logger.Warn("UpdateWebConfig Save db err %s", err.Error())
		return err.Error(), false
	}
	clearAllConfigCache()
	return "", true
}

func (dao *contentConfigDAO) Clean() {
	clearAllConfigCache()
	clearAllWriterCache()
	clearFrontAllCategoryCache()
	clearFrontAllBannerCache()
	clearFrontAllArticleCache()
	clearFrontAllMenuCache()
}
