package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var instance *sql.DB = newInstance()

func newInstance() *sql.DB {
	i, e := sql.Open("mysql", getUri())
	if e != nil {
		log.Fatal(e)
		return nil
	}
	return i
}

func GetInstance() *sql.DB {
	return instance
}

func getUri() string {

	if mustGetEnv("MODE") == "development" {
		return "root@tcp(127.0.0.1:3306)/cocomeshi"
	}

	dbUser := mustGetEnv("DB_USER")
	dbPwd := mustGetEnv("DB_PASS")
	instanceConnectionName := mustGetEnv("INSTANCE_CONNECTION_NAME")
	dbName := mustGetEnv("DB_NAME")

	return fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if k == "" {
		log.Printf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
