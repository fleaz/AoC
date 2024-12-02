package main

import "testing"

type testpair struct {
	mass     int
	wantFuel int
}

func TestCalculateFuel(t *testing.T) {
	pairs := []testpair{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, p := range pairs {
		isFuel := calculateFuel(p.mass)
		if isFuel != p.wantFuel {
			t.Errorf("Wanted %d fuel for mass %d but got %d", p.wantFuel, p.mass, isFuel)
		}

	}

}

func TestComplexCalculateFuel(t *testing.T) {
	pairs := []testpair{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, p := range pairs {
		isFuel := complexCalculateFuel(p.mass)
		if isFuel != p.wantFuel {
			t.Errorf("Wanted %d fuel for mass %d but got %d", p.wantFuel, p.mass, isFuel)
		}

	}

}
