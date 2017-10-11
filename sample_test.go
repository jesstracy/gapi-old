package main

import "testing"

func TestAdd(t *testing.T) {
	a := 3
	b := 4
	if x := Add(a, b); x != 7 {
		t.Errorf("Add(%v, %v) = %v, want %v", a, b, x, 7)
	}
}