package main

import "testing"

func TestArea_Triangle(t *testing.T) {

	t.Run("triangle", func(t *testing.T) {

		Atriangle := Triangle{4.0, 5.0}
		got := Atriangle.Area()
		want := 10.0

		if got != want {
			t.Errorf("got %g want %g,", got, want)
		}

	})

}

func TestPerimeter_Triangle(t *testing.T) {

	t.Run("triangle", func(t *testing.T) {

		Ptriangle := Triangle{4.0, 5.0}
		got := Ptriangle.Parimeter()
		want := 15.403124237432849

		if got != want {
			t.Errorf("got %g want %g,", got, want)
		}
	})
}
