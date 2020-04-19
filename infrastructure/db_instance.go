package infrastructure

import (
	"database/sql"
	"log"
)

var instance *sql.DB = newInstance()

func newInstance() *sql.DB {
	i, e := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cocomeshi")
	if e != nil {
		log.Fatal(e)
		return nil
	}
	return i
}

func GetInstance() *sql.DB {
	return instance
}
