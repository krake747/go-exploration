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
	catalog := map[int]bookstore.Book{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
		3: {Id: 3, Title: "Spark Joy"},
	}
	want := map[int]bookstore.Book{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
		3: {Id: 3, Title: "Spark Joy"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
		3: {Id: 3, Title: "Spark Joy"},
	}
	want := bookstore.Book{Id: 2, Title: "The Power of Go: Tools"}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
