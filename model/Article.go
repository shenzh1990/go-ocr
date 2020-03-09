package model

import "time"

type Article struct {
	Id        int64     `xorm:"pk autoincr comment('文章') INT(10)"`
	TagId     int64     `xorm:" notnull comment('标签ID') INT(10)"`
	Title     string    `xorm:"default '' comment('文章标题') VARCHAR(50)"`
	Desc      string    `xorm:"default '' comment('描述') VARCHAR(255)"`
	Content   string    `xorm:"TEXT comment('内容')"`
	CreatedBy string    `xorm:"default '' comment('创建人') VARCHAR(100)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	UpdatedBy string    `xorm:"default '' comment('修改人') VARCHAR(100)"`
	State     int       `xorm:"default 1 comment('状态 0为禁用1为启用') TINYINT(3)"`
}
