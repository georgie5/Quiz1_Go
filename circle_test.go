package main

import "testing"

func TestArea_Circle(t *testing.T) {

	t.Run("circle", func(t *testing.T) {
		Acircle := Circle{6}
		got := Acircle.Area()
		want := 113.09733552923255

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
func TestPermimeter_Circle(t *testing.T) {

	t.Run("circle", func(t *testing.T) {
		Pcircle := Circle{6}
		got := Pcircle.Parimeter()
		want := 37.69911184307752

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
