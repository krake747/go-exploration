package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Substract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
