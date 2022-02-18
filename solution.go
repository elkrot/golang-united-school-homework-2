package square

import (
"math"
)

type intCustomType int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3.0)) / 4.0
	case SidesSquare:
		return 4.0 * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0.0
	}
}