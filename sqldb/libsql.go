package sqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/echo")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
