package main

import(
  "fmt"

  "github.com/zeidlitz/dbserver/internal/env"
  "github.com/zeidlitz/dbserver/internal/server"
)

type Config struct{
  serverAddress string
  httpPort int
}

func main() {
  var cfg Config
  cfg.serverAddress = env.GetString("SERVER_ADDRESS", "localhost")
  cfg.httpPort = env.GetInt("HTTP_PORT", 8080)
  address := fmt.Sprintf(cfg.serverAddress + ":" + "%d", cfg.httpPort)
  server.Start(address)
}
