package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

//PostgresqlDataSource holds the information used to build the psql datasource
type PostgresqlDataSource struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

//newPostgresqlDataSource builds the new sqllite datasource
func newPostgresqlDataSource() (DataSource, error) {
	var pgDatasource = &PostgresqlDataSource{}
	var ok bool

	pgDatasource.Host, ok = os.LookupEnv("PGHOST")
	if !ok {
		pgDatasource.Host = "localhost"
	}

	pgDatasource.Port, ok = os.LookupEnv("PGPORT")
	if !ok {
		pgDatasource.Port = "5432"
	}

	pgDatasource.User, ok = os.LookupEnv("PGUSER")
	if !ok {
		pgDatasource.User = "postgres"
	}

	pgDatasource.Password, ok = os.LookupEnv("PGPASSWORD")
	if !ok {
		pgDatasource.Password = "postgres"
	}

	pgDatasource.Database, ok = os.LookupEnv("PGDATABASE")
	if !ok {
		pgDatasource.Database = "postgres"
	}

	pgDatasource.SSLMode, ok = os.LookupEnv("PGSSLMODE")
	if !ok {
		pgDatasource.SSLMode = "disable"
	}

	return pgDatasource, nil
}

//InitDB creates the initial connection to the database and create the application tables
// implements DataSource
func (p PostgresqlDataSource) InitDB() (*bun.DB, error) {
	ctx := context.Background()

	c := pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", p.Host, p.Port)),
		pgdriver.WithUser(p.User),
		pgdriver.WithPassword(p.Password),
		pgdriver.WithDatabase(p.Database),
	)
	if p.SSLMode != "disable" {
		panic(fmt.Sprintf("SSL mode not yet supported"))
	} else {
		c.Config().TLSConfig = nil
	}
	sqlDB := sql.OpenDB(c)

	//Check Connection is live
	if err := sqlDB.Ping(); err != nil {
		log.Errorf("Error connecting to DB %s", err)
		return nil, err
	}

	//Create the Database
	db := bun.NewDB(sqlDB, pgdialect.New())
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
