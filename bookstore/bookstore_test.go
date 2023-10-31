package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "The Hobbit",
		Author: "J.R. R. Tolkien",
		Copies: 2,
	}
}
