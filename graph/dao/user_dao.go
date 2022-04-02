package dao

import (
	"context"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

//InsertUser inserts model.user in to DB
func InsertUser(ctx context.Context, db *bun.DB, user *model.User) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for user %v insert,%s", user, err)
		return err
	}

	//Do insert record
	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Errorf("Error inserting user %s", err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing user %v,%s", user, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//UpdateUser updates model.user in to DB
func UpdateUser(ctx context.Context, db *bun.DB, user *model.User) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for user %v update,%s", user, err)
		return err
	}

	//Do insert record
	if _, err := tx.NewUpdate().Model(user).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error updating user %s", err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing user %v,%s", user, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//DeleteUser delete the user from the database
func DeleteUser(ctx context.Context, db *bun.DB, user *model.User) error {
	if _, err := db.NewDelete().Model(user).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error deleting todo %d,%s", user.ID, err)
		return err
	}

	return nil
}

//SelectUsers selects all users from the database
func SelectUsers(ctx context.Context, db *bun.DB, users *[]*model.User, last *int) error {
	var err error
	if last != nil {
		err = db.NewSelect().Model(users).Limit(*last).Scan(ctx)
	} else {
		err = db.NewSelect().Model(users).Scan(ctx)
	}

	log.Debugf("Users %v", users)

	if err != nil {
		log.Errorf("Error querying all users, %s", err)
		return err
	}

	return nil
}

//FindUserByID find the user by its primary key
func FindUserByID(ctx context.Context, db *bun.DB, user *model.User) error {
	if err := db.NewSelect().Model(user).WherePK().Scan(ctx); err != nil {
		log.Errorf("Error getting user with id %d, %s", user.ID, err)
		return err
	}

	return nil
}
