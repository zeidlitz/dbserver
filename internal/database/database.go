package database

type Database interface {
	Query(query string) (err error, response string)
  Connect(connection string) (err error)
}
