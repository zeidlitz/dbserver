package server

import(
  "net/http"
  "log/slog"
  "encoding/json"
)

type Database interface {
    Query(query string) (err error, response string)
}

var database Database

type Data struct {
	Response  string `json:"value"`
}

type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
  slog.Info("Handleded request" , "Remote Address", r.RemoteAddr)
  response := Response{
    Status: http.StatusOK,
    Message: "Running",
  }

  res, err := json.Marshal(response)
  if err != nil{
    http.Error(w,"Internal Server Error", http.StatusInternalServerError)
  }
  w.Write(res)
}

func queryHandler(w http.ResponseWriter, r *http.Request){
  slog.Info("Handleded request" , "Remote Address", r.RemoteAddr)
  // Temporary mockup query, refactor this
  query := "GET TRASH"
  err, queryResponse := database.Query(query)
  if err != nil {
    slog.Error("Error during query", "query", query, "error", err)
  }

  data := Data{
    Response: queryResponse,
  }

  response := Response{
    Status: http.StatusOK,
    Message: "Running",
    Data: data,
  }

  res, err := json.Marshal(response)
  if err != nil{
    http.Error(w,"Internal Server Error", http.StatusInternalServerError)
  }

  w.Write(res)
}

func Start(address string, db Database) {
  slog.Info("Starting up", "address", address)
  database = db
  http.HandleFunc("/", handler)
  http.HandleFunc("/query", queryHandler)
  err := http.ListenAndServe(address, nil)
  if err != nil {
    slog.Error("Could not startup", "error", err.Error())
    return
  }
}
