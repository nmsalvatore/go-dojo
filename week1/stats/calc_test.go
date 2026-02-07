package main

import "testing"

func TestMean(t *testing.T) {
	var meanTests = []struct {
		name    string
		in      []int
		out     float64
		wantErr bool
	}{
		{"multiple numbers", []int{5, 3, 8, 1, 4}, 4.2, false},
		{"one number", []int{5}, 5, false},
		{"includes negative number", []int{-1, 4, 6}, 3, false},
		{"zeros", []int{0, 0, 0}, 0, false},
		{"empty", []int{}, 0, true},
	}

	for _, tt := range meanTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mean(tt.in)

			if tt.wantErr && err == nil {
				t.Error("want error, didn't get one")
				return
			}

			want := tt.out
			if got != want {
				t.Errorf("got %.2f, want %.2f", got, want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	var medianTests = []struct {
		name    string
		in      []int
		out     float64
		wantErr bool
	}{
		{"odd count", []int{5, 3, 8, 1, 4}, 4.0, false},
		{"even count", []int{5, 3, 8, 1}, 4.0, false},
		{"one number", []int{5}, 5.0, false},
		{"includes negative number", []int{-15, 15}, 0.0, false},
		{"zeros", []int{0, 0, 0}, 0, false},
		{"empty", []int{}, 0, true},
	}

	for _, tt := range medianTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Median(tt.in)

			if tt.wantErr && err == nil {
				t.Error("want error, didn't get one")
				return
			}

			want := tt.out
			if got != want {
				t.Errorf("got %.2f, want %.2f", got, want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	var sumTests = []struct {
		name string
		in   []int
		out  int
	}{
		{"multiple numbers", []int{5, 4, 8, 1, 4}, 22},
		{"one number", []int{36}, 36},
		{"includes negative number", []int{3, -9}, -6},
		{"zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range sumTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.in)
			want := tt.out
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
