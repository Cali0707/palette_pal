package models

import (
	"fmt"
	"github.com/Cali0707/palette_pal/pkg/crypto"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	UserName  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func CreateUser(user User, db *sqlx.DB) (err error) {
	hashedPassword, err := crypto.HashPassword(user.Password)
	if err != nil {
		fmt.Printf("Unable to hash password!")
		return err
	}
	tx, err := db.Begin()
	_, err = tx.Exec("INSERT INTO users (username, email, password) values($1, $2, $3)", user.UserName, user.Email, hashedPassword)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	err = tx.Commit()
	return err
}

func GetUserByUsername(username string, db *sqlx.DB) (user User, err error) {
	var users []User
	err = db.Select(&users, "SELECT * FROM users where username=$1 LIMIT 1", username)
	return users[0], err
}
