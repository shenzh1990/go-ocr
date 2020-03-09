package model

import "time"

//CREATE TABLE `blog_tag` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(100) DEFAULT '' COMMENT '标签名称',
//`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

type Tag struct {
	Id        int64     `xorm:"pk autoincr comment('用户名') INT(10)"`
	Name      string    `json:"name" xorm:" notnull comment('标签名称') VARCHAR(100)"`
	CreatedBy string    `json:"created_by" xorm:"default '' comment('创建人') VARCHAR(100)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	UpdatedBy string    `json:"modified_by" xorm:"default '' comment('修改人') VARCHAR(100)"`
	State     int       ` json:"state" xorm:"default 1 comment('状态 0为禁用1为启用') TINYINT(3)"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	Db.Where(maps).Limit(pageSize, pageNum).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) int64 {
	count, _ := Db.Where(maps).Count(&Tag{})
	return count
}

func ExistTagByName(name string) (bool, error) {

	flag, err := Db.Where("name = ?", name).Exist(&Tag{})
	if err != nil {
		return false, err
	}

	return flag, nil
}

func AddTag(name string, state int, createdBy string) bool {
	Db.Insert(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}
func ExistTagByID(id int) bool {
	flag, _ := Db.Where("id = ?", id).Exist(&Tag{})
	return flag
}

func DeleteTag(id int) int64 {
	flag, err := Db.ID(id).Delete(&Tag{})
	if err != nil {
		return 0
	}
	return flag
}

func EditTag(id int, data interface{}) int64 {
	flag, err := Db.ID(id).Update(data)
	if err != nil {
		return 0
	}
	return flag
}
