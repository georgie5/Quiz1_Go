package main

import "testing"

func TestArea(t *testing.T) {
	area, _ := Square(8.0)
	got := area
	want := 64.0

	if got != want {
		t.Errorf("got %g want %g,S", got, want)
	}
}

func TestParimeter(t *testing.T) {
	_, perimeter := Square(8.0)
	got := perimeter
	want := 32.0

	if got != want {
		t.Errorf("got %g want %g,", got, want)
	}
}
