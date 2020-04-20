package model

import (
	"BitCoin/pkg/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var Db *xorm.Engine

func init() {
	var err error
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	Db, err = xorm.NewEngine(settings.BitConfig.Db.DriverName, settings.BitConfig.Db.DBUrl)
	if err != nil {
		fmt.Println(err)
	}
	Db.ShowSQL(false)
	Db.CreateTables(GymInfo{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}
