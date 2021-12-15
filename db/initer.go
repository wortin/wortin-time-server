package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wortin-time-server/conf"
	"wortin-time-server/po"
)

var DB *gorm.DB

func init() {
	var err error
	user := conf.Vip.GetString("mysql.user")
	pd := conf.Vip.GetString("mysql.password")
	host := conf.Vip.GetString("mysql.host")
	port := conf.Vip.GetString("mysql.port")
	db := conf.Vip.GetString("mysql.database")
	args := conf.Vip.GetString("mysql.args")

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pd, host, port, db, args)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 建表
	err = DB.AutoMigrate(&po.Tag{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&po.Action{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&po.ActionTag{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&po.Project{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&po.SubAction{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&po.Target{})
	if err != nil {
		panic(err)
	}

	// 初始化默认数据
	InsertDefaultTags()
}
