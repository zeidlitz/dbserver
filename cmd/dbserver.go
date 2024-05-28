package main

import (
	"fmt"
	"log/slog"

	"github.com/zeidlitz/dbserver/internal/env"
	"github.com/zeidlitz/dbserver/internal/server"
	"github.com/zeidlitz/dbserver/internal/databasefactory"
)

type Config struct {
	serverAddress string
	httpPort      int
	dbConnection  string
	dbType  string
}

func main() {
	var cfg Config

	cfg.httpPort = env.GetInt("SERVER_PORT", 8080)
	cfg.serverAddress = env.GetString("SERVER_ADDRESS", "localhost")
	cfg.dbType = env.GetString("DB_TYPE", "sqlite")
	cfg.dbConnection = env.GetString("DB_CONNECTION", "database.db")
	address := fmt.Sprintf(cfg.serverAddress+":"+"%d", cfg.httpPort)
  err, db := databasefactory.GetDatabase(cfg.dbType, cfg.dbConnection)
  if err != nil {
    slog.Error("Error creating database", "error", err)
    panic(err)
  }
	server.Start(address, db)
}
