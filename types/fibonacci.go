package types

import "math/big"

type Fibonacci struct {
	N uint16   `json:"n"`
	V *big.Int `json:"v"`
}
