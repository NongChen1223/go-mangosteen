package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysqlHost     = "172.18.0.2"
	mysqlPort     = 3306
	mysqlUser     = "mangosteen"
	mysqlPassword = "123456"
	mysqlDbname   = "mangosteen_dev"
)

func MysqlConnect() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln("sql链接失败 打印", err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln("sql ping失败 打印:", err)
	}
	log.Println("sql链接成功")
}
func MysqlClose() {
	DB.Close()
	log.Println("sql链接关闭")
}
