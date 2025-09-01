package main

import "testing"

func TestAddTwoNumber(t *testing.T) {
	got := AddTwoNumber(10, 20)
	want := 30

	if got != want {
		t.Errorf("AddTwoNumber(10, 30) = %d; want %d ", got, want)
	}

}

func TestAddTwoNumberTableriven(t *testing.T) {
	tests := []struct {
		a, b, s int
	}{
		{1, 2, 3},
		{100, 200, 300},
		{400, 500, 900},
	}

	for _, tt := range tests {
		got := AddTwoNumber(tt.a, tt.b)
		want := tt.s

		if got != want {
			t.Errorf("AddTwoNumber(10, 30) = %d; want %d ", got, want)
		}
	}
}
