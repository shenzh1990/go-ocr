package model

import (
	"time"
)

type Rule struct {
	Id           int64     `xorm:"pk autoincr"`
	Text         string    `xorm:"index   notnull comment('名称')  VARCHAR(128)"`
	Frequency    int       `xorm:"notnull comment('频次') INT(10)"`
	PartOfSpeech string    `xorm:" comment('词性') VARCHAR(2) "`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

func (r *Rule) GetAllRules() (rules []Rule) {
	Db.Find(&rules)
	return
}

func (r *Rule) AddRule() bool {
	_, err := Db.Insert(r)
	if err != nil {
		return false
	}
	return true
}
