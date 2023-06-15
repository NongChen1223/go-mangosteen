package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysqlHost     = "localhost"
	mysqlPort     = 3307
	mysqlUser     = "mangosteen"
	mysqlPassword = "123456"
	mysqlDbname   = "mangosteen_dev"
)

func MysqlConnect() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDbname)
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
func MysqlCreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users(
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	email VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`
	_, err := DB.Exec(createTableSQL)

	if err != nil {
		log.Fatalln("创建表失败 打印:", err)
	}

	log.Println("创建表成功")
}

func MysqlClose() {
	DB.Close()
	log.Println("sql链接关闭")
}
