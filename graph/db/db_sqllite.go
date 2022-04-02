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
	"path"
)

//SqlliteDataSource holds the information used to build the sqllite datasource
type SqlliteDataSource struct {
	DbFilePath string
}

//newSqlliteDataSource builds the new sqllite datasource
func newSqlliteDataSource() (DataSource, error) {
	dbFilePath, ok := os.LookupEnv("TODOS_DB_FILE")
	if !ok {
		cwd, err := os.Getwd()
		if err != nil {
			log.Errorf("Error getting current working directory %s", err)
			return nil, err
		}

		err = os.MkdirAll(path.Join(cwd, "work"), os.ModeDir)
		if err != nil {
			log.Errorf("Error making db directory %s", err)
			return nil, err
		}

		dbFilePath = fmt.Sprintf("%s/todo.db", path.Join(cwd, "work"))
	}
	return &SqlliteDataSource{DbFilePath: dbFilePath}, nil
}

//InitDB creates the initial connection to the database and create the application tables
// implements DataSource
func (d SqlliteDataSource) InitDB() (*bun.DB, error) {
	ctx := context.Background()
	sqlDB, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cached=shared", d.DbFilePath))
	if err != nil {
		log.Errorf("Error connecting to DB %s", err)
		return nil, err
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
		if _, err := db.NewDropTable().IfExists().Model((*model.User)(nil)).Exec(ctx); err != nil {
			log.Errorf("Unable to drop table user %s", err)
		}
		if _, err := db.NewDropTable().IfExists().Model((*model.Todo)(nil)).Exec(ctx); err != nil {
			log.Errorf("Unable to drop table todo %s", err)
		}
	}

	if _, err := db.NewCreateTable().IfNotExists().Model((*model.User)(nil)).Exec(ctx); err != nil {
		return nil, err
	}
	if _, err := db.NewCreateTable().IfNotExists().Model((*model.Todo)(nil)).Exec(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
