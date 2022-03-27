package fcbrs

import (
	"testing"
)

func TestCalcMove(t *testing.T) {
	tests := []struct {
		Name   string
		Values []int
		Move   int
		Want   []int
	}{
		{
			Name:   "Move 0",
			Values: []int{8, 8, 8},
			Move:   0,
			Want:   []int{11, 8, 7},
		},
		{
			Name:   "Move 1",
			Values: []int{8, 8, 8},
			Move:   1,
			Want:   []int{5, 6, 11},
		},
		{
			Name:   "Move 2",
			Values: []int{8, 8, 8},
			Move:   2,
			Want:   []int{10, 8, 6},
		},
		{
			Name:   "Move 3",
			Values: []int{8, 8, 8},
			Move:   3,
			Want:   []int{6, 10, 8},
		},
		{
			Name:   "Move 4",
			Values: []int{8, 8, 8},
			Move:   4,
			Want:   []int{8, 5, 10},
		},
		{
			Name:   "Move 5",
			Values: []int{8, 8, 8},
			Move:   5,
			Want:   []int{8, 11, 6},
		},
		{
			Name:   "Max",
			Values: []int{8, 15, 8},
			Move:   5,
			Want:   []int{8, 16, 6},
		},
		{
			Name:   "Min",
			Values: []int{8, 8, 2},
			Move:   5,
			Want:   []int{8, 11, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := calcMove(tt.Values, tt.Move)

			if out[0] != tt.Want[0] || out[1] != tt.Want[1] || out[2] != tt.Want[2] {
				t.Errorf("Got: %v - Want: %v", out, tt.Want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		Name string
		A    int
		B    int
		C    int
		Want []string
	}{
		{
			Name: "simple test",
			A:    5,
			B:    8,
			C:    9,
			Want: []string{"Investment", "Investment", "Blackmail", "Donation", "Orgy"},
		},
		{
			Name: "non perfect moveset",
			A:    9,
			B:    7,
			C:    11,
			Want: []string{"Blackmail", "Donation", "Donation", "Orgy", "Teambuilding"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			out := Solve(tt.A, tt.B, tt.C)
			if len(out) == 0 {
				t.Error("no perfect moveset found")
			}
			if len(out) != 5 {
				t.Error("output slice has wrong size")
			}
			if out[0] != tt.Want[0] || out[1] != tt.Want[1] || out[2] != tt.Want[2] || out[3] != tt.Want[3] || out[4] != tt.Want[4] {
				t.Errorf("Got: %v - Want: %v", out, tt.Want)
			}
		})
	}
}
