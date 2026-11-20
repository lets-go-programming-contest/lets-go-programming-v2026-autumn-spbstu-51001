package temperature

import (
	"errors"
	"fmt"
)

const (
	MinTemperature = 15
	MaxTemperature = 30
	Unreachable    = -1

	opGreaterOrEqual = ">="
	opLessOrEqual    = "<="
)

var ErrInvalidOperator = errors.New("operator must be \">=\" or \"<=\"")

type Range struct {
	lower int
	upper int
}

func NewRange() Range {
	return Range{
		lower: MinTemperature,
		upper: MaxTemperature,
	}
}

func (r Range) Apply(operator string, value int) (Range, error) {
	switch operator {
	case opGreaterOrEqual:
		r.lower = max(r.lower, value)
	case opLessOrEqual:
		r.upper = min(r.upper, value)
	default:
		return r, fmt.Errorf("%w, got %q", ErrInvalidOperator, operator)
	}

	return r, nil
}

func (r Range) Optimal() int {
	if r.lower > r.upper {
		return Unreachable
	}

	return r.lower
}
