package postgres

import "database/sql"

type Database struct {
	db *sql.DB
}
