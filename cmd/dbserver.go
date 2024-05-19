package main

import (
	"fmt"

	"github.com/zeidlitz/dbserver/internal/env"
	"github.com/zeidlitz/dbserver/internal/server"
	"github.com/zeidlitz/dbserver/internal/sqlite"
)

type Config struct {
	serverAddress string
	httpPort      int
	dbConnection  string
}

func main() {
	var cfg Config

	cfg.serverAddress = env.GetString("SERVER_ADDRESS", "localhost")
	cfg.httpPort = env.GetInt("HTTP_PORT", 8080)
	cfg.dbConnection = env.GetString("DB_TYPE", "sqlite")
	cfg.dbConnection = env.GetString("DB_CONNECTION", "database.db")
	address := fmt.Sprintf(cfg.serverAddress+":"+"%d", cfg.httpPort)
  db := sqlite.SQLite{Connection: cfg.dbConnection}
	server.Start(address, db)
}
