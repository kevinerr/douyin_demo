package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, //缓存预编译语句
	})
	//db.LogMode(true)
	if err != nil {
		panic(err)
	}
	//if gin.Mode() == "release" {
	//	db.LogMode(false)
	//}
	//db.SingularTable(true)       //默认不加复数s
	//db.DB().SetMaxIdleConns(20)  //设置连接池，空闲
	//db.DB().SetMaxOpenConns(100) //打开
	//db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//migration()
}
