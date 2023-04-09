package ex22

import (
	"fmt"
	"math/big"
)

// Since the condition didn't say anything about our own implementation of big number addition,
// I'll just use the big package. I don't see the point in reinventing the wheel

func Test() {
	a := big.NewFloat(1214532748961243189533092075893274980748730927413290743297.90248378904132704891274902317893427891032497132074198)
	b := big.NewFloat(21986756369778248472109837489125646201387478139741789432654327186489.23456324352643263426346325432698)
	fmt.Printf("[main] a + b = %g\n", a.Add(a, b))
	fmt.Printf("[main] a - b = %g\n", a.Sub(a, b))
	fmt.Printf("[main] a * b = %g\n", a.Mul(a, b))
	fmt.Printf("[main] a / b = %g\n", a.Quo(a, b))
}
