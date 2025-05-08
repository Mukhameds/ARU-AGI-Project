package main

import (
	"fmt"
	"math/cmplx"
	
)

func zetaApprox(s complex128, nMax int) complex128 {
	var sum complex128
	for n := 1; n <= nMax; n++ {
		ns := cmplx.Pow(complex(float64(n), 0), s)
		sum += 1 / ns
	}
	return sum
}

func main() {
	nMax := 1000
	step := 0.1
	threshold := 1.0 // Порог для нуля
	fmt.Println("Возможные нули ζ(0.5 + i·t):")

	for t := -50.0; t <= 50.0; t += step {
		s := complex(0.5, t)
		absZ := cmplx.Abs(zetaApprox(s, nMax))

		// Локальный минимум и значение близко к 0
		s1 := complex(0.5, t-step)
		s2 := complex(0.5, t+step)
		absPrev := cmplx.Abs(zetaApprox(s1, nMax))
		absNext := cmplx.Abs(zetaApprox(s2, nMax))

		if absZ < threshold && absZ < absPrev && absZ < absNext {
			fmt.Printf("t ≈ %.3f  |ζ(s)| ≈ %.4f\n", t, absZ)
		}
	}
}
