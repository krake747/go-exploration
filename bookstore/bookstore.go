package bookstore

import "errors"

// Book represents information about a book.
type Book struct {
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
