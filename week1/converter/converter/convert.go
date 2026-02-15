package converter

import (
	"math"
)

func Convert(val float64, from, to string) (float64, error) {
	source, ok := units[from]
	if !ok {
		return 0, &UnknownUnitError{from}
	}

	target, ok := units[to]
	if !ok {
		return 0, &UnknownUnitError{to}
	}

	if source.Category != target.Category {
		return 0, &IncompatibleUnitsError{from, to}
	}

	base := source.ToBase(val)
	result := target.FromBase(base)
	rounded := math.Round(result*100) / 100

	return rounded, nil
}
