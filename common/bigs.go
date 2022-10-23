package common

import (
    "math/big"
)

var (
    BIG_ZERO = new(big.Int)
    BIG_ONE = new(big.Int).SetUint64(1)
)

// a greater then b
func BigIntCmpGt(a *big.Int, b *big.Int) bool {
    return new(big.Int).Set(a).Cmp(b) > 0
}

// a greater than or equals b
func BigIntCmpGte(a *big.Int, b *big.Int) bool {
    return new(big.Int).Set(a).Cmp(b) >= 0
}

// a less than b
func BigIntCmpLt(a *big.Int, b *big.Int) bool {
    return new(big.Int).Set(a).Cmp(b) < 0
}

// a less than or equals b
func BigIntCmpLte(a *big.Int, b *big.Int) bool {
    return a.Cmp(b) <= 0
}
