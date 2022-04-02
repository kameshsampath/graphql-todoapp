package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

//InitDB creates the initial connection to the database and create the application tables
func InitDB(dbFilePath string) (*bun.DB, error) {
	sqlDB, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cached=shared", dbFilePath))
	if err != nil {
		log.Errorf("Error connecting to DB %s", err)
	}

	//Create the Database
	db := bun.NewDB(sqlDB, sqlitedialect.New())
	// BUNDEBUG=1 logs failed queries
	// BUNDEBUG=2 logs all queries
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.FromEnv("BUNDEBUG")))
	//Create Tables
	//BUN_SCHEMA_GEN_MODE - drop-and-create or update
	bunSchemGenMode := os.Getenv("BUN_SCHEMA_GEN_MODE")
	if bunSchemGenMode == "drop-and-create" {
		if _, err := db.NewDropTable().IfExists().Model((*model.User)(nil)).Exec(context.TODO()); err != nil {
			log.Errorf("Unable to drop table user %s", err)
		}
		if _, err := db.NewDropTable().IfExists().Model((*model.Todo)(nil)).Exec(context.TODO()); err != nil {
			log.Errorf("Unable to drop table todo %s", err)
		}
		if _, err := db.NewCreateTable().Model((*model.User)(nil)).Exec(context.TODO()); err != nil {
			return nil, err
		}
		if _, err := db.NewCreateTable().Model((*model.Todo)(nil)).Exec(context.TODO()); err != nil {
			return nil, err
		}
	}

	return db, err
}
