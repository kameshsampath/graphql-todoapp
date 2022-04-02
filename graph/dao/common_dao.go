package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

//Insert is generic function to insert data of type T
func Insert[T any](ctx context.Context, db *bun.DB, t *T) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for %v insert,%s", t, err)
		return err
	}

	//Do insert record
	if _, err := tx.NewInsert().Model(t).Exec(ctx); err != nil {
		log.Errorf("Error inserting %v %s", t, err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing %v,%s", t, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//Update is generic function to update data of type T
func Update[T any](ctx context.Context, db *bun.DB, t *T) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for %v insert,%s", t, err)
		return err
	}

	//Do Update record
	if _, err := tx.NewUpdate().Model(t).Exec(ctx); err != nil {
		log.Errorf("Error inserting %v %s", t, err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing %v,%s", t, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//Delete is generic function to delete data of type T
func Delete[T any](ctx context.Context, db *bun.DB, t *T) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for %v insert,%s", t, err)
		return err
	}

	//Do Delete record
	if _, err := tx.NewDelete().Model(t).Exec(ctx); err != nil {
		log.Errorf("Error inserting %v %s", t, err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing %v,%s", t, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//SelectAll selects all record of type T from the database
func SelectAll[T any](ctx context.Context, db *bun.DB, ms *[]*T, last *int) error {
	var err error
	if last != nil {
		err = db.NewSelect().Model(ms).Limit(*last).Scan(ctx)
	} else {
		err = db.NewSelect().Model(ms).Scan(ctx)
	}

	log.Debugf("Got records %v", ms)

	if err != nil {
		log.Errorf("Error querying all records, %s", err)
		return err
	}

	return nil
}

//FindUserByPrimaryKey find the record by its primary key
func FindUserByPrimaryKey[T any](ctx context.Context, db *bun.DB, m *T) error {
	if err := db.NewSelect().Model(m).WherePK().Scan(ctx); err != nil {
		log.Errorf("Error getting record by primary key, %s", err)
		return err
	}

	return nil
}
