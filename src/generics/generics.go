package generics

import (
	"fmt"
	"io"
)

func PrintAnything[T any](v T) {
	fmt.Println(v)
}

func Identity[T any](v T) T {
	return v
}

func PrintAnythingTo[T any](w io.Writer, v T) {
	fmt.Fprintln(w, v)
}
