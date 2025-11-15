package utils

import (
	"errors"
	"math/big"
	"strings"
)

// ParseEthToWei converts "2.5" ETH into 2500000000000000000 wei.
func ParseEthToWei(s string) (*big.Int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty string")
	}

	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return nil, errors.New("invalid ETH amount")
	}

	wei := new(big.Int)
	_, ok := wei.SetString(parts[0], 10)
	if !ok {
		return nil, errors.New("invalid integer part")
	}

	wei.Mul(wei, big.NewInt(1_000_000_000_000_000_000)) // * 1e18

	// Decimal part
	if len(parts) == 2 {
		decimals := parts[1]
		if len(decimals) > 18 {
			return nil, errors.New("too many decimal places")
		}

		decimals = decimals + strings.Repeat("0", 18-len(decimals))

		frac := new(big.Int)
		_, ok := frac.SetString(decimals, 10)
		if !ok {
			return nil, errors.New("invalid decimal part")
		}

		wei.Add(wei, frac)
	}

	return wei, nil
}
