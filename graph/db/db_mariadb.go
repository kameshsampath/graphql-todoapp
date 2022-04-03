package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

//MariadbDataSource holds the information used to build the mysql or mariadb datasource
type MariadbDataSource struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

//newMariadbDataSource builds the new mysql or mariadb datasource
func newMariadbDataSource() (DataSource, error) {
	var mariadbDataSource = &MariadbDataSource{}
	var ok bool

	mariadbDataSource.Host, ok = os.LookupEnv("MARIADB_HOST")
	if !ok {
		mariadbDataSource.Host = "localhost"
	}

	mariadbDataSource.Port, ok = os.LookupEnv("MARIADB_PORT")
	if !ok {
		mariadbDataSource.Port = "3306"
	}

	mariadbDataSource.User, ok = os.LookupEnv("MARIADB_USER")
	if !ok {
		mariadbDataSource.User = "root"
	}

	mariadbDataSource.Password, ok = os.LookupEnv("MARIADB_PASSWORD")
	if !ok {
		mariadbDataSource.Password, ok = os.LookupEnv("MARIADB_ROOT_PASSWORD")
		if !ok {
			mariadbDataSource.Password = "password"
		}
	}

	mariadbDataSource.Database, ok = os.LookupEnv("MARIADB_DATABASE")
	if !ok {
		mariadbDataSource.Database = "demodb"
	}

	return mariadbDataSource, nil
}

//InitDB creates the initial connection to the database and create the application tables
// implements DataSource
func (m MariadbDataSource) InitDB() (*bun.DB, error) {
	ctx := context.Background()

	c := mysql.Config{
		Addr:                 fmt.Sprintf("%s:%s", m.Host, m.Port),
		User:                 m.User,
		Passwd:               m.Password,
		DBName:               m.Database,
		AllowNativePasswords: true,
	}

	log.Printf("DSN:%s", c.FormatDSN())
	sqlDB, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Errorf("Error connecting to DB %s", err)
		return nil, err
	}

	//Create the Database
	db := bun.NewDB(sqlDB, mysqldialect.New())
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
