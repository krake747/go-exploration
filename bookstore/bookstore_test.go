package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "The Hobbit",
		Author: "J.R. R. Tolkien",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "The Hobbit",
		Author: "J.R. R. Tolkien",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "The Hobbit",
		Author: "J.R. R. Tolkien",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Id: 1, Title: "For the Love of Go"},
	}
	want := bookstore.Book{Id: 1, Title: "For the Love of Go"}
	got := bookstore.GetBook(catalog, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
