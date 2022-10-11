package hw8

import (
	"testing"
)

func Test1(t *testing.T) {
	N := 10
	a := 3
	got := Pollard(N,b)
	want := "Composite"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}
