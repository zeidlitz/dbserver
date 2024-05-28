package trashdatabase

import (
	"log/slog"
)

type TrashDB struct {
	Name string
}

func (t TrashDB) Connect(connection string) error {
	slog.Info("TrashDB connected", "connection", connection)
	return nil
}

func (t TrashDB) Query(query string) (err error, response string) {
	slog.Info("Recieved query", "query", query)
	return nil, "TRASH RESPONSE"
}
