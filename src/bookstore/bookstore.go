package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

// Category represents information about a book's category.
type Category int

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
	category        Category
}

func (book Book) NetPriceCents() int {
	return book.PriceCents - (book.PriceCents * book.DiscountPercent / 100)
}

func (book *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	book.PriceCents = price
	return nil
}

func (book *Book) SetCategory(category Category) error {
	var validCategory = map[Category]bool{
		CategoryAutobiography:     true,
		CategoryLargePrintRomance: true,
		CategoryParticlePhysics:   true,
	}
	if !validCategory[category] {
		return fmt.Errorf("unknown category %q", category)
	}
	book.category = category
	return nil
}

func (book Book) Category() Category {
	return book.category
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
type Catalog map[int]Book
