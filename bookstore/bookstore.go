package bookstore

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
