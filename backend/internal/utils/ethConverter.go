package utils

import (
	"errors"
	"math/big"

	"github.com/shopspring/decimal"
)

func ParseUSDC(amount string) (*big.Int, error) {
	dec, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	dec = dec.Mul(decimal.NewFromInt(1_000_000))

	result := new(big.Int)
	result, ok := result.SetString(dec.StringFixed(0), 10)
	if !ok {
		return nil, errors.New("failed to convert decimal to big.Int")
	}

	return result, nil
}
