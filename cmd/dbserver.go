package main

import(
  "fmt"

  "github.com/zeidlitz/dbserver/internal/env"
  "github.com/zeidlitz/dbserver/internal/server"
  "github.com/zeidlitz/dbserver/internal/trashdatabase"
)


type Config struct{
  serverAddress string
  httpPort int
  dbConnection string
}

func main() {
  var cfg Config
  var db trashdatabase.TrashDB

  cfg.serverAddress = env.GetString("SERVER_ADDRESS", "localhost")
  cfg.httpPort = env.GetInt("HTTP_PORT", 8080)
  cfg.dbConnection = env.GetString("DB_CONNECTION", "database.db")
  address := fmt.Sprintf(cfg.serverAddress + ":" + "%d", cfg.httpPort)
  server.Start(address, db)
}
