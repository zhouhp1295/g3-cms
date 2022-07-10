package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
)

type contentTagDAO struct {
	crud.BaseDao
}

var ContentTagDao = &contentTagDAO{
	crud.BaseDao{Model: new(model.ContentTag)},
}

func (dao *contentTagDAO) InsertOrGetByTitle(title string) *model.ContentTag {
	cnt := dao.CountByColumn("title", title)
	m := new(model.ContentTag)
	if cnt > 0 {
		//已存在，取并返回
		if crud.DbSess().Where("title = ? and deleted = ?", title, crud.FlagNo).First(m).Error == nil {
			return m
		}
	} else {
		//新的标签
		m.Title = title
		if crud.DbSess().Create(m).Error == nil {
			return m
		}
	}
	return nil
}

func (dao *contentTagDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentTag); _ok {
		cnt := dao.CountByColumn("title", _m.Title)
		if cnt > 0 {
			msg = "已存在标签名"
			return
		}
		ok = true
	}
	return
}

type FrontTagData struct {
	Id    interface{} `json:"id"`
	Title interface{} `json:"title"`
	Cnt   interface{} `json:"cnt"`
}

func (dao *contentTagDAO) FrontTags() []FrontTagData {
	tagRows := make([]map[string]interface{}, 0)

	crud.DbSess().Model(new(model.ContentArticleTag)).
		Joins("left join content_tag on content_article_tag.tag_id = content_tag.id").
		Where("content_tag.deleted = ? and content_tag.status = ?", crud.FlagNo, crud.FlagYes).
		Select("content_tag.id,content_tag.title,count(content_tag.id) cnt").
		Group("content_tag.id").
		Order("cnt desc").
		Limit(20).
		Find(&tagRows)

	result := make([]FrontTagData, len(tagRows))

	for i, row := range tagRows {
		result[i].Id = row["id"]
		result[i].Title = row["title"]
		result[i].Cnt = row["cnt"]
	}

	return result
}
