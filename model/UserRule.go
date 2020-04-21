package model

import (
	"time"
)

type UserRule struct {
	Id           int64     `xorm:"pk autoincr"`
	Text         string    `xorm:"index   notnull comment('名称')  VARCHAR(128)"`
	Frequency    int       `xorm:"notnull comment('频次') INT(10)"`
	PartOfSpeech string    `xorm:" comment('词性') VARCHAR(2) "`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

func (r *UserRule) GetAllUserRules() (rules []UserRule) {
	Db.Find(&rules)
	return
}
func (r *UserRule) GetUserRules() (rules []UserRule) {
	Db.Where("text=?", r.Text).Find(&rules)
	return
}
func (r *UserRule) AddUSerRule() bool {
	_, err := Db.Insert(r)
	if err != nil {
		return false
	}
	return true
}
