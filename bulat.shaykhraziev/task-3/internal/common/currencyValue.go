package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CurrencyValue uint64

const (
	scale       = 4
	scaleFactor = 10_000
)

var ErrNegativeValue = errors.New("negative numbers cannot be CurrencyValue")

func (cv *CurrencyValue) UnmarshalText(b []byte) error {
	text := strings.TrimSpace(string(b))
	if strings.HasPrefix(text, "-") {
		return ErrNegativeValue
	}

	intPart, fracPart, _ := strings.Cut(text, ",")

	if len(fracPart) > scale {
		fracPart = fracPart[:scale]
	}

	fracPart += strings.Repeat("0", scale-len(fracPart))

	n, err := strconv.ParseUint(intPart+fracPart, 10, 64)
	if err != nil {
		return fmt.Errorf("parse currency value %q: %w", text, err)
	}

	*cv = CurrencyValue(n)

	return nil
}

func (cv CurrencyValue) MarshalJSON() ([]byte, error) {
	v := uint64(cv)

	return fmt.Appendf(nil, "%d.%0*d", v/scaleFactor, scale, v%scaleFactor), nil
}
