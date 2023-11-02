package generics_test

import (
	"bytes"
	"errors"
	"fmt"
	"generics"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ExamplePrintAnything() {
	generics.PrintAnything[int](5)
	// Output:
	// 5
}

func ExampleIdentity() {
	i := generics.Identity[string]("Gopher")
	err := errors.New("Gopher! What did you do?")
	fmt.Println(i)
	fmt.Println(generics.Identity(err))
	// Output:
	// Gopher
	// Gopher! What did you do?
}

func TestPrintAnythingTo_PrintsInputToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	generics.PrintAnythingTo[string](buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
