package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	biga := big.NewInt(int64(1 << 22))
	bigb := big.NewInt(int64(1 << 27))

	fmt.Println(biga)
	fmt.Println(bigb)

	fmt.Println("+++ ", add(biga, bigb))

	fmt.Println("---", sub(biga, bigb))

	fmt.Println("***", mul(biga, bigb))

	fmt.Println("///", div(bigb, biga))
}
