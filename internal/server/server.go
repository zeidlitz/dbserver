package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/zeidlitz/dbserver/internal/database"
)

type Database = database.Database

var db Database

type Data struct {
	Response string `json:"value"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handleded request", "Remote Address", r.RemoteAddr)
	response := Response{
		Status:  http.StatusOK,
		Message: "Running",
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Write(res)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handleded request", "Remote Address", r.RemoteAddr)
	query := "SELECT * FROM key_value_pairs"
	err, queryResponse := db.Query(query)
	if err != nil {
		slog.Error("Error during query", "query", query, "error", err)
	}

	data := Data{
		Response: queryResponse,
	}

	response := Response{
		Status:  http.StatusOK,
		Message: "Running",
		Data:    data,
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write(res)
}

func Start(address string, database Database) {
	slog.Info("Starting up", "address", address)
	db = database
	http.HandleFunc("/", handler)
	http.HandleFunc("/query", queryHandler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		slog.Error("Could not startup", "error", err.Error())
		return
	}
}
