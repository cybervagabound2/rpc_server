package model

import (
	"database/sql"
	"github.com/pkg/errors"
)

type user struct {
	UUID        int	   `json:"uuid"`
	Username    string `json:"username"`
	Registered  string `json:"registered"`
}

func (u *user) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *user) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *user) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}