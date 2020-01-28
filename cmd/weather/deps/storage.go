package deps

import (
	"database/sql"

	"github.com/danilvpetrov/weathersource/storage/mysqlstorage"
)

// ProvideStorage provides *mysqlstorage.Storage.
func ProvideStorage(
	db *sql.DB,
) *mysqlstorage.Storage {
	return &mysqlstorage.Storage{
		DB: db,
	}
}
