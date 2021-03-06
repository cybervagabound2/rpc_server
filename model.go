package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	ID    	    int    `json:"id"`
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	Registered  string `json:"registered"`
}

// get user
func (u *user) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT uuid, username, registered FROM users WHERE id=$1",
		u.ID).Scan(&u.UUID, &u.Username, &u.Registered)
}

// update user
func (u *user) updateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET username=$1 WHERE id=$2",
			u.Username, u.ID)

	return err
}

// delete user
func (u *user) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}

// create user
func (u *user) createUser(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO users(uuid, username) VALUES($1, $2) RETURNING id",
		genUuid(), u.Username).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

// fetch list of users
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	rows, err := db.Query(
		"SELECT id, uuid, username, registered FROM users LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.UUID, &u.Username, &u.Registered); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// function to generate unique uuid for each user
func genUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
