package db

import (
	"fmt"
	"github.com/uptrace/bun"
	"os"
)

//DataSource provide the method(s) to set up the database and its tables
type DataSource interface {
	InitDB() (*bun.DB, error)
}

//NewDB gets builds the right datasource based on the type
func NewDB() (DataSource, error) {
	dbType := os.Getenv("TODO_DB_TYPE")
	switch dbType {
	case "postgresql":
		panic(fmt.Sprintf("postgresql not yet supported"))
	case "mysql":
		panic(fmt.Sprintf("mysql not yet supported"))
	default:
		return newSqlliteDataSource()
	}
}
