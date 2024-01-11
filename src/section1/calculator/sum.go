package calculator

// private
var offset float64 = 1

// public
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	return a + b + offset
}
