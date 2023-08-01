package worker

import "math/big"

var fibs = make(map[uint16]*big.Int)

func Work(n uint16) *big.Int {
	switch n {
	case 0:
		return big.NewInt(0)
	case 1:
		return big.NewInt(1)
	}

	if fib, found := fibs[n]; found {
		return fib
	}

	fibs[n] = new(big.Int).Add(Work(n-1), Work(n-2))
	return fibs[n]
}
