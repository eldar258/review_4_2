package migration

import (
	"github.com/jmoiron/sqlx"
)

const createCities = `CREATE TABLE IF NOT EXISTS cities ( id SERIAL PRIMARY KEY, name VARCHAR(30) NOT NULL, state VARCHAR(30) NOT NULL )`

func Migrate(db *sqlx.DB) error {
	_, err := db.Exec(createCities)
	return err
}
