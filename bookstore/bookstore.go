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

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, id int) Book {
	for _, b := range catalog {
		if b.Id == id {
			return b
		}
	}
	return Book{}
}
