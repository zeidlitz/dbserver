package databasefactory

import (
	"errors"
	"log/slog"

	"github.com/zeidlitz/dbserver/internal/database"
	"github.com/zeidlitz/dbserver/internal/sqlite"
	"github.com/zeidlitz/dbserver/internal/trashdatabase"
)

type Database = database.Database


func GetDatabase(dbtype string, dbconnection string) (err error, db Database){
  switch dbtype {
  case "sqlite":
	  db = sqlite.SQLite{Connection: dbconnection}
    err = nil
  case "trashdatabse":
    db = trashdatabase.TrashDB{Name: "trashconnection"}
    err = nil
  default:
    slog.Error("Undefined database type", "dbtype", dbtype)
    err = errors.New("Undefined database type")
  }
  return err, db
}
