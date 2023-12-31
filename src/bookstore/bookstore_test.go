package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "The Hobbit",
		Author: "J.R. R. Tolkien",
		Copies: 2,
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{Title: "For the Love of Go", PriceCents: 4000, DiscountPercent: 25}
	want := 3000
	if got := book.NetPriceCents(); want != got {
		t.Errorf("want %d, got %d", want, got)
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
	if got := result.Copies; want != got {
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
	if _, err := bookstore.Buy(b); err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "The Power of Go: Tools"},
		3: {Id: 3, Title: "Spark Joy"},
	}
	want := []bookstore.Book{
		{Id: 1, Title: "For the Love of Go"},
		{Id: 2, Title: "The Power of Go: Tools"},
		{Id: 3, Title: "Spark Joy"},
	}
	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].Id < got[j].Id
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
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
	got, err := bookstore.GetBook(catalog, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIdReturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	if _, err := bookstore.GetBook(catalog, 999); err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:      "For the love of Go",
		PriceCents: 4000,
	}
	want := 3000
	if err := book.SetPriceCents(want); err != nil {
		t.Fatal(err)
	}
	if got := book.PriceCents; want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	if err := book.SetPriceCents(-1); err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title: "For the Love of Go",
	}
	categories := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, category := range categories {
		if err := book.SetCategory(category); err != nil {
			t.Fatal(err)
		}
		if got := book.Category(); category != got {
			t.Errorf("want category %q, got %q", category, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title: "For the Love of Go",
	}
	if err := book.SetCategory(999); err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}
