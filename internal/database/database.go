package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

func Query(query string) error {
	db, err := sql.Open("sqlite3", "pkg/storage/database.db")

	if err != nil {
		slog.Error("Error when opening databse", "Error", err.Error())
		return err
	}

	defer db.Close()

  slog.Info("Running query" ,"query", query)
	_, err = db.Exec(query)

	if err != nil {
		slog.Error("Error when doing INSERT: ", err)
		return err
	}

	return err
}

