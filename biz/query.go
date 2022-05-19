package biz

import (
	"context"
	"gentest/dal"
	"gentest/dal/model"
	"gentest/dal/query"
)

var q = query.Use(dal.DB.Debug())

func Create(ctx context.Context)  {
	var err error
	ud := q.User.WithContext(ctx)

	userData := &model.User{Name: "peter", Age: 22}
	err = ud.Create(userData)
	catchErr("create one item fail", err)
}