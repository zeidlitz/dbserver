package main

import(
  "fmt"
  "log/slog"

  "github.com/zeidlitz/kvdb/internal/env"
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
  slog.Info("Starting up", "address", address)
}
