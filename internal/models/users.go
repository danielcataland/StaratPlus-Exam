package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	db "stratplusapi/internal/database"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
}

func (u User) CreateUsr() error {
	cnn := db.GetConnection()

	log.Println("[INFO]: Creating user")
	query := `INSERT INTO users (email, password, username, phone) 
			VALUES(?, ?, ?, ?)`

	stmt, err := cnn.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	qresult, err := stmt.Exec(u.Email, getMD5Hash(u.Password), u.UserName, u.Phone)
	if err != nil {
		return err
	}

	if i, err := qresult.RowsAffected(); err != nil || i != 1 {
		return errors.New("[ERROR]: Rows not afected ")
	}
	log.Println("[INFO]: User added to DB")
	return nil
}

func (u *User) GetUser(us User, login bool) (User, string) {
	cnn := db.GetConnection()

	query := ``

	if login {
		query = `SELECT * FROM users WHERE 
				(email=? OR username=?) && password=? LIMIT 1`

		rows, err := cnn.Query(query, us.Email, us.UserName, getMD5Hash(us.Password))

		for rows.Next() {
			rows.Scan(&u.Email, &u.Password, &u.UserName, &u.Phone)
		}

		if err != nil {
			log.Println("[INFO] Error to execute query: ", err)
			return User{}, "Internal server error"
		}
	} else {
		log.Println("[INFO]: Get user without login")
		query = `SELECT * FROM users WHERE email=? OR phone=? LIMIT 1`

		rows, err := cnn.Query(query, us.Email, us.Phone)

		for rows.Next() {
			rows.Scan(&u.Email, &u.Password, &u.UserName, &u.Phone)
		}

		if err != nil {
			log.Println("[INFO] Error to execute query: ", err)
			return User{}, "Internal server error"
		}
	}

	log.Println("[INFO]: Email exist: ", u.Email)
	return *u, "exist"
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
