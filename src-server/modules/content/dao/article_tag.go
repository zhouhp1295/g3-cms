// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
)

type contentArticleTagDAO struct {
	crud.BaseDao
}

var ContentArticleTagDao = &contentTagDAO{
	crud.BaseDao{Model: new(model.ContentArticleTag)},
}
