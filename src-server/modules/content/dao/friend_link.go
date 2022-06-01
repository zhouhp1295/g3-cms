package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
)

type contentFriendLinkDAO struct {
	crud.BaseDao
}

var ContentFriendLinkDao = &contentFriendLinkDAO{
	crud.BaseDao{Model: new(model.ContentFriendLink)},
}
