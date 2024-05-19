package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
  Connection string
	conn *sql.DB
}

func (s SQLite) Connect() (err error) {
  slog.Info("Connecting", "connection", s.Connection)
  s.conn, err = sql.Open("sqlite3", s.Connection)

	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		return err
	}
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

  cols, err := rows.Columns() 
  values := make([]interface{}, len(cols))
  for i := range cols {
      values[i] = new(string)
  }

	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			slog.Error("Error scanning rows", "rows", rows, "error", err.Error())
			return err, ""
		}
	}

  // value := strings.Join(values, " ")
  value := fmt.Sprintf("%v", &values)
	return err, value
}
