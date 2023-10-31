package main

import "fmt"

func main() {
	title := "For the Love of Go"
	author := "John Arundel"
	copies := 99
	inStock := true
	royaltyPercentage := 12.5
	isSpecialOffer := true

	printTitle(title)
	fmt.Println("Author: ", author)
	fmt.Println("Copies: ", copies)
	fmt.Println("Stock: ", inStock)
	fmt.Println("Royalty %: ", royaltyPercentage)
	fmt.Println("Special Offer: ", isSpecialOffer)
}

func printTitle(title string) {
	fmt.Println("Title: ", title)
}
