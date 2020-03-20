package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewInstance() (*sql.DB, error) {
	return sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cocomeshi")
}
