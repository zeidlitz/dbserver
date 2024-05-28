package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Connection string
	conn       *sql.DB
}

func (s SQLite) Connect(connection string) (err error) {
	s.Connection = connection
	return nil
}

func (s SQLite) Query(query string) (err error, response string) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		slog.Error("Error when connecting")
		return err, ""
	}

	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		slog.Error("Error when executing", "query", query, "error", err.Error())
		return err, ""
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		slog.Error("Error when retriveing columns", "error", err.Error())
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			slog.Error("Error when scanning values", "error", err)
		}
	}

	for i, col := range columns {
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			fmt.Printf("%s: %s\n", col, string(b))
		} else {
			fmt.Printf("%s: %v\n", col, val)
		}
	}

	return err, ""
}
