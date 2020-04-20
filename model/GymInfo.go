package model

import "time"

type GymInfo struct {
	Id   int64  `xorm:"pk autoincr comment('id') INT(10)"`
	Name string `json:"name" xorm:" notnull comment('
体育馆名称') VARCHAR(100)"`
	Lng string `json:"lng" xorm:"  comment('
经度') VARCHAR(128)"`
	Lat string `json:"lat" xorm:"  comment('
 维度') VARCHAR(128)"`
	Address string `json:"address" xorm:" notnull comment('
 维度') VARCHAR(500)"`
	CreatedBy string    `json:"created_by" xorm:"default '' comment('创建人') VARCHAR(100)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	UpdatedBy string    `json:"modified_by" xorm:"default '' comment('修改人') VARCHAR(100)"`
}

func (c *GymInfo) GetGymInfos(pageNum int, pageSize int, maps interface{}) (GymInfos []GymInfo) {
	Db.Where(maps).Limit(pageSize, pageNum).Find(&GymInfos)
	return
}

func (c *GymInfo) GetAllGymInfos() (GymInfos []GymInfo) {
	Db.Find(&GymInfos)
	return
}
func (c *GymInfo) AddGymInfo() bool {
	Db.Insert(c)
	return true
}
func (c *GymInfo) UpdateGymInfo(maps interface{}) bool {
	Db.Table(c).Update(maps)
	return true
}
