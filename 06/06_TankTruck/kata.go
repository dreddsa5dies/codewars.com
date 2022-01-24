/*
Caution: This kata does not currently have any known supported versions for Go. It may not be completable due to dependencies on out-dated libraries/language versions.
To introduce the problem think to my neighbor who drives a tanker truck. The level indicator is down and he is worried because he does not know if he will be able to make deliveries. We put the truck on a horizontal ground and measured the height of the liquid in the tank.

Fortunately the tank is a perfect cylinder and the vertical walls on each end are flat. The height of the remaining liquid is h, the diameter of the cylinder base is d, the total volume is vt (h, d, vt are positive or null integers). You can assume that h <= d.

Could you calculate the remaining volume of the liquid? Your function tankvol(h, d, vt) returns an integer which is the truncated result (e.g floor) of your float calculation.

Examples:
tankvol(40,120,3500) should return 1021 (calculation gives about: 1021.26992027)

tankvol(60,120,3500) should return 1750

tankvol(80,120,3500) should return 2478 (calculation gives about: 2478.73007973)
*/

package kata

import "math"

func TankVol(h, d, vt int) int {
	v := 0
	pi := math.Pi
	length := (4 * float64(vt)) / (float64(d) * float64(d) * pi)

	if 2*h <= d {
		theta := math.Acos((float64(d) - 2*float64(h)) / float64(d))
		smax := 0.25 * float64(d) * float64(d) * theta
		smin := 0.125 * float64(d) * float64(d) * math.Sin(2*theta)
		v = int((smax - smin) * length)
	} else { // 2*h > d
		theta := math.Acos((float64(2*h) - float64(d)) / float64(d))
		smax := 0.25 * float64(d) * float64(d) * theta
		smin := 0.125 * float64(d) * float64(d) * math.Sin(2*theta)
		v = int(float64(vt) - (smax-smin)*length)
	}

	return v
}
