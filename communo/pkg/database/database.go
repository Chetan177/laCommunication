package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

type Details struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	database   = "lacommunicationdb"
	connString = "chetan:chetan@tcp(127.0.0.1:3306)/"
	table      = "details"
)

/*
Table Structure

CREATE TABLE `lacommunicationdb`.`details` ( `email` VARCHAR(100) NOT NULL , `password` VARCHAR(50) NOT NULL , PRIMARY KEY (`email`));


Conn :
<user>:<pass>@tcp(127.0.0.1:3306)/<Db>
*/

func InsertData(email, pass string) error {

	db, err := sql.Open("mysql", connString+database)

	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s VALUES ( '%s', '%s' )", table, email, pass)
	insert, err := db.Query(query)
	if err != nil {
		return err
	}

	defer insert.Close()
	defer db.Close()

	return err
}

func GetData(email string) (string, string, error) {

	db, err := sql.Open("mysql", connString+database)
	if err != nil {
		return "", "", err
	}

	var details Details
	query := fmt.Sprintf("SELECT email, password  FROM %s where email = ?",table)
	err = db.QueryRow(query, email).Scan(&details.Email, &details.Password)
	if err != nil {
		return "", "", err
	}

	defer db.Close()

	return details.Email, details.Password,err
}
