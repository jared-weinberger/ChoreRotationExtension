package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}
type UserStore struct {
	db *sqlx.DB
}

func MakeUserStore(dbFilePath string) (*UserStore, error) {
	db, err := sqlx.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(INIT_USERS_TABLE)
	if err != nil {
		return nil, err
	}
	return &UserStore{db: db}, nil
}

func (store *UserStore) GetUser(email string) (*user, error) {
	var user user
	err := store.db.Select(&user, SELECT_USER_BY_EMAIL, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *UserStore) PutUser(email string, name string) error {
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(PUT_USER)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
