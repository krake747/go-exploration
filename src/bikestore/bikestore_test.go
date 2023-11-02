package bikestore_test

import (
	"bikestore"
	"fmt"
	"testing"
)

func TestCustomer(t *testing.T) {
	t.Parallel()
	_ = bikestore.Customer{
		Name: "Gopher",
	}
}

func ExampleCustomer() {
	c := bikestore.Customer{
		Name: "Gopher",
	}
	fmt.Println(c)
	// Output:
	// {Gopher}
}
