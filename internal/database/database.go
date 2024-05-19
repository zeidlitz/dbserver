package database

type Database interface {
	Query(query string) (err error, response string)
	Connect() (err error)
}
