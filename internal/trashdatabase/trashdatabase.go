package trashdatabase

import(
  "log/slog"
)

type TrashDB struct {
    name string
}

func (t TrashDB) Query(query string) (err error, response string)  {
  slog.Info("Recieved query", "query", query)
  return nil, "TRASH RESPONSE"
}
