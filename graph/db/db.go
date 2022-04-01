package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gocloud.dev/postgres"
)

//ConnectToDB creates the initial connection to the database
func ConnectToDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))
	db, err := postgres.Open(context.TODO(), connStr)
	if err != nil {
		log.Errorf("Error connecting to DB %s", err)
	}
	return db, err
}
