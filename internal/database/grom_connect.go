package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
)

const (
	mysqlHost   = "localhost"
	mysqlDbname = "mangosteen_dev"
)

// gormConnect gorm链接
func gormConnect() {
	dsn := fmt.Sprintf("user:pass@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlHost, mysqlDbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // 表名前缀，`User`表为`t_users`
			SingularTable: true,                              // 使用单数表名，启用该选项，此时，`User` 表将是`user`
			NoLowerCase:   true,                              // 表名不转小写
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
	if err != nil {
		log.Fatalln("gorm连接数据库失败 打印", err)
	}
}
