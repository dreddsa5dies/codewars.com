package kata

import (
	"math"
)

func RectangleRotation(a, b int) int {
	aHalfBisect := (float64(a) / math.Sqrt(2)) / 2
	bHalfBisect := (float64(b) / math.Sqrt(2)) / 2

	rect1 := []float64{round(aHalfBisect)*2 + 1, round(bHalfBisect)*2 + 1}
	rect2 := []float64{}

	if (aHalfBisect - round(aHalfBisect)) < 0.5 {
		rect2 = append(rect2, rect1[0]-1)
	} else {
		rect2 = append(rect2, rect1[0]+1)
	}

	if (bHalfBisect - round(bHalfBisect)) < 0.5 {
		rect2 = append(rect2, rect1[1]-1)
	} else {
		rect2 = append(rect2, rect1[1]+1)
	}

	return int(rect1[0]*rect1[1] + rect2[0]*rect2[1])
}

func round(x float64) float64 {
	if math.IsNaN(x) {
		return x
	}
	if x == 0.0 {
		return x
	}
	roundFn := math.Ceil
	if math.Signbit(x) {
		roundFn = math.Floor
	}
	xOrig := x
	x -= math.Copysign(0.5, x)
	if x == 0 || math.Signbit(x) != math.Signbit(xOrig) {
		return math.Copysign(0.0, xOrig)
	}
	if x == xOrig-math.Copysign(1.0, x) {
		return xOrig
	}
	r := roundFn(x)
	if r != x {
		return r
	}
	return roundFn(x*0.5) * 2.0
}
