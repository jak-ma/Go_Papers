package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Drive() *sql.DB {
	var dns = "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
