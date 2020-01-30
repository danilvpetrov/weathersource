package deps

import (
	"database/sql"

	// This blank import is required to initialize MySQL database connection as
	// the dependency.
	_ "github.com/go-sql-driver/mysql"

	"github.com/danilvpetrov/weathersource/internal/envvar"
)

// ProvideDatabase provides an instance of MySQL database connection, the caller
// is responsible for closing the connection.
func ProvideDatabase() (*sql.DB, func(), error) {
	dsn, err := envvar.StringRequired("WEATHER_DB_DSN")
	if err != nil {
		return nil, nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}

	close := func() {
		db.Close()
	}

	return db, close, nil
}
