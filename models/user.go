package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/go-api/db"
	"example.com/go-api/utils"
)

type User struct {
	Id        int64
	Email     string `binding: "required"`
	Password  string `binding: "required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user User) Save() error {
	query := `INSERT INTO users(email, password, created_at, updated_at) values (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Print("insert failed")
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Print("password hash error")
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword, time.Now().UTC(), time.Now().UTC())

	if err != nil {
		fmt.Print("insert exec failed", err.Error())
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		fmt.Print("iLastInsertId")
		return err
	}

	return err
}

func (user *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, user.Email)
	var retievedPassword string
	err := row.Scan(&user.Id, &retievedPassword)

	if err != nil {
		fmt.Print("password scan failed")
		return err
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, retievedPassword)

	if !isValidPassword {
		fmt.Print("password not valid")
		return errors.New("credentials invalid")
	}

	return nil
}
