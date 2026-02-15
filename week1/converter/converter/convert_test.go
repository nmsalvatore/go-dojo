package converter

import (
	"errors"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	var successTests = []struct {
		name string
		val  float64
		from string
		to   string
		want float64
	}{
		{"km to ft", 5, "kilometers", "feet", 16_404.20},
		{"m to yd", 34, "meters", "yards", 37.18},
		{"mi to ft", 8, "miles", "feet", 42239.9},
		{"kg to g", 10, "kilograms", "grams", 10_000},
		{"lb to g", 1, "pounds", "grams", 453.6},
		{"oz to kg", 600, "ounces", "kilograms", 17.01},
		{"F to K", 100, "fahrenheit", "kelvin", 310.93},
		{"F to C", 90, "fahrenheit", "celsius", 32.22},
		{"C to F", 0, "celsius", "fahrenheit", 32},
		{"same unit", 120, "miles", "miles", 120},
		{"large value", 40_000_000, "miles", "kilometers", 64_373_600},
		{"zero value", 0, "meters", "kilometers", 0},
		{"negative value", -5, "kilometers", "meters", -5_000},
		{"fractional", 0.5, "miles", "feet", 2_639.99},
	}

	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.val, tt.from, tt.to)
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}

	var unknownUnitTests = []struct {
		name string
		val  float64
		from string
		to   string
	}{
		{"single unknown unit", 200, "dudes", "kilometers"},
		{"multiple unknown units", 400, "dudes", "bros"},
	}

	for _, tt := range unknownUnitTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Convert(tt.val, tt.from, tt.to)
			if err == nil {
				t.Error("wanted error, didn't get one")
			}

			var uue *UnknownUnitError
			if err != nil && !errors.As(err, &uue) {
				t.Error("wanted UnknownUnitError, got something else")
			}

			var msg = "unknown unit"
			if !strings.Contains(err.Error(), msg) {
				t.Errorf("%q missing from error message %q", msg, err.Error())
			}
		})
	}

	var incompatibleUnitTests = []struct {
		name string
		val  float64
		from string
		to   string
	}{
		{"kg to ft", 10, "kilograms", "feet"},
		{"g to m", 10, "grams", "meters"},
		{"lb to mi", 10, "pounds", "miles"},
		{"oz to yd", 10, "ounces", "yards"},
		{"F to yd", 10, "fahrenheit", "yards"},
		{"g to C", 10, "grams", "celsius"},
		{"ft to K", 10, "feet", "kelvin"},
	}

	for _, tt := range incompatibleUnitTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Convert(tt.val, tt.from, tt.to)
			if err == nil {
				t.Error("wanted error, didn't get one")
			}

			var iue *IncompatibleUnitsError
			if err != nil && !errors.As(err, &iue) {
				t.Error("wanted IncompatibleUnitsError, got something else")
			}

			var msg = "incompatible units"
			if !strings.Contains(err.Error(), msg) {
				t.Errorf("%q missing from error message %q", msg, err.Error())
			}
		})
	}
}
