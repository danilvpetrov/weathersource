package mysqlstorage

import (
	"database/sql"
)

// Storage is MySQL storage implementation.
type Storage struct {
	DB *sql.DB
}
