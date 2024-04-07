package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"go-ocr/pkg/settings"
)

var Db *xorm.Engine

func Start() {
	var err error
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	Db, err = xorm.NewEngine(settings.OcrConfig.Db.DriverName, settings.OcrConfig.Db.DBUrl)
	if err != nil {
		fmt.Println(err)
	}
	Db.ShowSQL(false)
	//Db.CreateTables(UserRule{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}
