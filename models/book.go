package models

import (
	"go-school-library/database"
	"time"
)

func GetBookID(b Book) int {
	return b.id
}

func GetBookName(id int) string {
	var book Book
	db, err := database.GetDB()
	if err != nil {
		return err.Error()
	}
	bookRow, err := db.Query("SELECT name FROM books WHERE id = ?", id)
	for bookRow.Next() {

		err = bookRow.Scan(&book.name)
		if err != nil {
			return err.Error()
		}
	}
	return book.name
}

func GetBookDescription(b Book) string {
	return b.name
}

func CreateBook(name string, description string) error {
	db, err := database.GetDB()
	database.CreateBooksTable()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		"INSERT INTO books (name, description) VALUES (?, ?)",
		name,
		description,
	)
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

func DeleteBook(id int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

func RentBook(accountId int, bookId int) error {
	delivery_date := time.Now().
		AddDate(0, 0, 15).
		Format("02/01")
		// This adds 15 days to deliver and formats to DD/MM

	err := database.CreateAccountBooksTable()
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	accountName := GetAccountName(accountId)
	bookName := GetBookName(bookId)

	_, err = db.Exec(
		"INSERT INTO accountBooks (accountID, accountName, bookID, bookName, deliveryDate) VALUES (?, ?, ?, ?, ?)",
		accountId,
		accountName,
		bookId,
		bookName,
		delivery_date,
	)

	return err
}

func ReturnBook(accountID int, bookID int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		"DELETE FROM accountBooks WHERE accountID = ? and bookID = ?",
		accountID,
		bookID,
	)
	return err
}
