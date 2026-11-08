package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type CurrencyValue uint64

const scale = 4

func (cv *CurrencyValue) UnmarshalText(b []byte) error {
	s := strings.TrimSpace(string(b))
	if strings.HasPrefix(s, "-") {
		return errors.New("negative numbers cannot be CurrencyValue")
	}

	intPart, fracPart, _ := strings.Cut(s, ",")

	fracPart = fracPart[:scale]
	fracPart += strings.Repeat("0", scale-len(fracPart))

	n, err := strconv.ParseUint(intPart+fracPart, 10, 64)
	if err != nil {
		return err
	}

	*cv = CurrencyValue(n)
	return nil
}

func (cv CurrencyValue) MarshalJSON() ([]byte, error) {
	sc := int64(math.Pow10(scale))
	v := int64(cv)
	return fmt.Appendf(nil, "%d.%04d", v/sc, v%sc), nil
}
