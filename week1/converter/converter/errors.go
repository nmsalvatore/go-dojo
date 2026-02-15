package converter

import (
	"fmt"
)

type UnknownUnitError struct {
	UnitName string
}

type IncompatibleUnitsError struct {
	FromUnitName string
	ToUnitName   string
}

func (e *UnknownUnitError) Error() string {
	return fmt.Sprintf("unknown unit %q", e.UnitName)
}

func (e *IncompatibleUnitsError) Error() string {
	return fmt.Sprintf("incompatible units %q and %q", e.FromUnitName, e.ToUnitName)
}
