package server

import(
  "net/http"
  "log/slog"
  "encoding/json"
)

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

func Start(address string){
  slog.Info("Starting up", "address", address)
  http.HandleFunc("/", handler)
  err := http.ListenAndServe(address, nil)
  if err != nil {
    slog.Error("Could not startup", "error", err.Error())
    return
  }
}
