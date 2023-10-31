package bookstore

import (
	"errors"
)

// Book represents information about a book.
type Book struct {
	Id     int
	Title  string
	Author string
	Copies int
}

// Customer represents information about a bookstore customer.
type Customer struct {
	Name  string
	Email string
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func GetAllBooks(catalog map[int]Book) map[int]Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) Book {
	return catalog[id]
}
