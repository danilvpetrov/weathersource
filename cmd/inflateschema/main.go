package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/danilvpetrov/weathersource/internal/envvar"
	"github.com/danilvpetrov/weathersource/storage/mysqlstorage"
)

var forceSchema bool

func init() {
	flag.BoolVar(
		&forceSchema,
		"force",
		false,
		"force overwriting the schema even if it already exists",
	)
}

func isSchemaExists(db *sql.DB, name string) (bool, error) {
	var exists bool
	// check if the schema already exists
	r := db.QueryRow(`SELECT EXISTS(
		SELECT SCHEMA_NAME FROM
			INFORMATION_SCHEMA.SCHEMATA
			WHERE SCHEMA_NAME = ?
		)`,
		mysqlstorage.SchemaName,
	)

	err := r.Scan(&exists)
	return exists, err
}

func run() error {
	flag.Parse()

	dsn, err := envvar.StringRequired("WEATHER_DB_DSN")
	if err != nil {
		return err
	}

	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return err
	}

	// It is critical to enable multistatements when running statements from SQL
	// file.
	if !cfg.MultiStatements {
		cfg.MultiStatements = true
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	if !forceSchema {
		ok, err := isSchemaExists(db, mysqlstorage.SchemaName)
		if err != nil {
			return err
		}

		if ok {
			fmt.Printf(
				"schema '%s' already exists at %s, skipping schema application...\n",
				mysqlstorage.SchemaName,
				cfg.Addr,
			)
			return nil
		}
	}

	q := string(mysqlstorage.MustAsset("schema.sql"))
	if _, err := db.Exec(q); err != nil {
		return err
	}

	fmt.Printf(
		"schema '%s' has been successfully applied against %s\n",
		mysqlstorage.SchemaName,
		cfg.Addr,
	)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
