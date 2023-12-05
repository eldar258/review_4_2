package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:my-secret-pw@(localhost:3306)/test")
}
