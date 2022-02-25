package models

import "time"

type BaseInfo struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	UpdateAt  time.Time `orm:"auto_now_add;type(datetime)" description:"更新时间"`
}
