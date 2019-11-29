package lib

import (
	m "math"
)

type Operation func(x float64, y float64) float64

type Opcion struct {
	X float64
	Y float64
	Operation Operation
}

func OptionOne(x float64, y float64) float64 {
	// 3𝑥^2 + 2𝑥𝑦 − 5𝑦^2
	return 3*(m.Pow(x, 2)) + (2*x*y) - (5*(m.Pow(y, 2)))
}

func OptionTwo(x float64, y float64) float64 {
	// (𝑥2 − 3𝑥 + 𝑦3)𝑠𝑒𝑛(|𝑥 − 𝑦|) − 𝑐𝑜𝑠(|𝑥 − 𝑦|)
	return (m.Pow(x, 2) - 3*x + m.Pow(y, 3)) *
			m.Sin(m.Abs(x - y)) -
			m.Cos(m.Abs(x - y))
}

func OptionThree(x float64, y float64) float64 {
	// −(𝑥 ∗ 𝑠𝑖𝑛^2(𝑥) ∗ 𝑐𝑜𝑠^3(𝑥)) + (𝑦 ∗ 𝑠𝑖𝑛^2(𝑦) ∗ 𝑐𝑜𝑠^3(𝑦))
	return -(x * m.Pow(m.Sin(x), 2) * m.Pow(m.Cos(x), 3)) +
			(y * m.Pow(m.Sin(y), 2) * m.Pow(m.Cos(y), 3))
}

func OptionFour(x float64, y float64) float64 {
	// |𝑠𝑖𝑛(𝑥) + 𝑐𝑜𝑠(𝑦) + √|𝑥 − 𝑦||
	return m.Abs(m.Sin(x) + m.Cos(y) + m.Sqrt(m.Abs(x - y)))
}