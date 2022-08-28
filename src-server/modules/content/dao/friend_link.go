// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

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
