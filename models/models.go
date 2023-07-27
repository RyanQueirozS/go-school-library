package models

import (
	"go-school-library/database"
)

type Book struct {
	id          int
	name        string
	description string
}

type Account struct {
	id   int
	name string
}

func CreateAccount(name string) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO accounts (name) VALUES (?)", name)
	return err
}

func DeleteAccount(id int64) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM accounts WHERE id = ?", id)
	return err
}

func GetAllAccounts() ([]Account, error) {
	accounts := []Account{}
	db, err := database.GetDB()
	if err != nil {
		return accounts, err
	}

	rows, err := db.Query("SELECT id, name FROM accounts")
	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		var account Account

		err = rows.Scan(&account.id, &account.name)
		if err != nil {
			return accounts, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func CreateBook(name string, description string) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO books (name, description) VALUES (?, ?)", name, description)
	return err
}

func GetAllBooks() ([]Book, error) {
	books := []Book{}
	db, err := database.GetDB()
	if err != nil {
		return books, err
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.id, &book.name, &book.description)
		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func DeleteBook(id int64) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
