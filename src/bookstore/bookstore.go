package bookstore

import (
	"errors"
	"fmt"
)

// Customer represents information about a bookstore customer.
type Customer struct {
	Name  string
	Email string
}

// Book represents information about a book.
type Book struct {
	Id              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

func (book Book) NetPriceCents() int {
	return book.PriceCents - (book.PriceCents * book.DiscountPercent / 100)
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func GetAllBooks(catalog Catalog) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func GetBook(catalog Catalog, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("id %d doesn't exist", id)
	}
	return b, nil
}

// Catalog represents information about a bookstore's collection of books.
type Catalog = map[int]Book
