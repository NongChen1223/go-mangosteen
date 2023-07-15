package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

const (
	gormMysqlUser     = "mangosteen"
	gormMysqlPassword = "123456"
	gormMysqlHost     = "localhost"
	gormMysqlPost     = "3307"
	gormMysqlDbname   = "mangosteen_dev"
)

// gormConnect gorm链接
func GormConnect() {
	/*
		DSN(Data Source Name)是数据源名称的缩写，它是一种按照一定格式描述数据库连接信息的字符串。
		DSN的格式如下：
		gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local
		用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=True&loc=Local
	*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", gormMysqlUser, gormMysqlPassword, gormMysqlHost, gormMysqlPost, gormMysqlDbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "mango_",                          // 表名前缀，`User`表为`t_users`
			SingularTable: true,                              // 使用单数表名，启用该选项，此时，`User` 表将是`user`
			NoLowerCase:   true,                              // 表名不转小写
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键 (代码里自动体现外键关系，数据库里不体现)
	})
	if err != nil {
		log.Fatalln("gorm连接数据库失败 打印", err)
	}

	// 获取通用数据库对象 sql.DB 以使用其函数
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("gorm获取数据库连接对象失败:", err)
	}
	// SetMaxIdleConns 设置空闲连接池的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置数据库打开连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	GormDB = db
	fmt.Println("gorm连接数据库成功", db)
	//type User struct {
	//	Name string
	//	Age  int
	//}
	//_ = db.AutoMigrate(&User{})

}
