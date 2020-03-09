package model

import (
	. "BitCoin/common"
	"time"
)

type Customer struct {
	Id        int64     `xorm:"pk autoincr"`
	Username  string    `xorm:"index unique notnull"`
	Password  string    `xorm:"notnull"`
	CreatedAt time.Time `xorm:"created"`
}

func (c *Customer) Register() string {
	Db.Insert(c)
	has, _ := Db.Get(c)
	if has {
		return JsonResponse(0, "注册成功")
	} else {
		return JsonResponse(1, "注册失败")
	}
}
func (c *Customer) Login() string {
	has, _ := Db.Get(c)
	if has {
		return JsonResponse(0, "登录成功")
	} else {
		return JsonResponse(1, "登录失败")
	}
}
func (c *Customer) SelectAllUser() string {
	var user []Customer
	Db.SQL("select * from customer").Find(&user)
	return JsonResponse(1, user)
}
func (c *Customer) SelectUserByName(username string) Customer {
	var user Customer
	Db.Where("username=?", username).Get(&user)
	return user
}
