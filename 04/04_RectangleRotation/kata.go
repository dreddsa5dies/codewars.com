package kata

import (
	"math"
)

func RectangleRotation(a, b int) int {
	pt := 0
	radius := math.Pow(float64(a)/2, 2) + math.Pow(float64(b)/2, 2)
	radius = round(math.Pow(radius, .5))

	mCos := math.Cos(float64(-45) * (math.Pi / 180.0))
	mSin := math.Sin(float64(-45) * (math.Pi / 180.0))

	for i := -radius; i <= radius+1; i++ {
		for j := -radius; j <= radius+1; j++ {
			x := i*mCos - j*mSin
			y := i*mSin + j*mCos
			if float64(-a/2) <= x && x <= float64(a/2) && float64(-b/2) <= y && y <= float64(b/2) {
				pt++
			}
		}
	}
	return pt
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
