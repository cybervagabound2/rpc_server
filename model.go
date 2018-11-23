package model

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
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

func (u *user) addUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}

// function to generate uuid for user
func (u *user) genUuid(db *sql.DB) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
