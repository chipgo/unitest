package add

import "math/big"

const (
	maxPrecInput uint = 1024
	roundPrec         = 5
)

func TwoBigFloat(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}

	x, ok := new(big.Float).SetPrec(maxPrecInput).SetString(a)
	if !ok {
		return ""
	}

	y, ok := new(big.Float).SetPrec(maxPrecInput).SetString(b)
	if !ok {
		return ""
	}

	total := new(big.Float).SetPrec(maxPrecInput).Add(x, y).
		Text('f', roundPrec)

	return total
}
