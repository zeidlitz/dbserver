package database

type Database interface {
    Query(dbConnection string, query string) (err error, response string)
}
