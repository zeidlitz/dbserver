package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

type SQLite struct {
	conn *sql.DB
}

func (s SQLite) Query(dbConnection string, query string) (err error, response string) {
	db, err := sql.Open("sqlite3", dbConnection)

	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		return err, ""
	}

	rows, err := db.Query(query)
	if err != nil {
		slog.Error("Error when executing", "query", query, "error", err.Error())
		return err, ""
	}

	defer rows.Close()

	var value string
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			slog.Error("Error scanning rows", "error", err.Error())
			return err, ""
		}
	}

	return err, value
}
