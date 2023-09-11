package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var _ = godotenv.Load(".env")
var (
	_   = godotenv.Load(".env")
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}

func CreateAccountBooksTable() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS accountBooks (accountID MEDIUMINT NOT NULL, accountName CHAR(50) NOT NULL, bookID MEDIUMINT NOT NULL, bookNAME CHAR(50) NOT NULL, deliveryDate CHAR(5) NOT NULL, PRIMARY KEY (accountID));",
	)
	return err
}
func CreateAccountsTable() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id MEDIUMINT NOT NULL AUTO_INCREMENT, name CHAR(50) NOT NULL, cpf CHAR(11), PRIMARY KEY(id));",
	)
	return err
}
func CreateBooksTable() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS books (id MEDIUMINT NOT NULL AUTO_INCREMENT, name CHAR(50) NOT NULL, description CHAR(150), PRIMARY KEY(id));",
	)

	return err
}
