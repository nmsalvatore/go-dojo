package converter

type UnitCategory string

const (
	Distance    UnitCategory = "distance"
	Mass        UnitCategory = "mass"
	Temperature UnitCategory = "temperature"
)

type Unit struct {
	Category UnitCategory
	ToBase   func(val float64) float64
	FromBase func(val float64) float64
}

var units = map[string]Unit{
	"feet": {
		Category: Distance,
		ToBase:   func(ft float64) float64 { return ft * 0.3048 },
		FromBase: func(m float64) float64 { return m / 0.3048 },
	},
	"kilometers": {
		Category: Distance,
		ToBase:   func(km float64) float64 { return km * 1000 },
		FromBase: func(m float64) float64 { return m / 1000 },
	},
	"meters": {
		Category: Distance,
		ToBase:   func(m float64) float64 { return m },
		FromBase: func(m float64) float64 { return m },
	},
	"miles": {
		Category: Distance,
		ToBase:   func(mi float64) float64 { return mi * 1609.34 },
		FromBase: func(m float64) float64 { return m / 1609.34 },
	},
	"yards": {
		Category: Distance,
		ToBase:   func(yd float64) float64 { return yd * 0.9144 },
		FromBase: func(m float64) float64 { return m / 0.9144 },
	},
	"kilograms": {
		Category: Mass,
		ToBase:   func(kg float64) float64 { return kg * 1000 },
		FromBase: func(g float64) float64 { return g / 1000 },
	},
	"grams": {
		Category: Mass,
		ToBase:   func(g float64) float64 { return g },
		FromBase: func(g float64) float64 { return g },
	},
	"pounds": {
		Category: Mass,
		ToBase:   func(lb float64) float64 { return lb * 453.6 },
		FromBase: func(g float64) float64 { return g / 453.6 },
	},
	"ounces": {
		Category: Mass,
		ToBase:   func(o float64) float64 { return o * 28.35 },
		FromBase: func(g float64) float64 { return g / 28.35 },
	},
	"kelvin": {
		Category: Temperature,
		ToBase:   func(k float64) float64 { return k },
		FromBase: func(k float64) float64 { return k },
	},
	"fahrenheit": {
		Category: Temperature,
		ToBase:   func(f float64) float64 { return (f + 459.67) / 1.8 },
		FromBase: func(k float64) float64 { return (k * 1.8) - 459.67 },
	},
	"celsius": {
		Category: Temperature,
		ToBase:   func(c float64) float64 { return c + 273.15 },
		FromBase: func(k float64) float64 { return k - 273.15 },
	},
}
