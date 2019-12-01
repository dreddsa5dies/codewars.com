/*
Write a function that takes a positive integer n, sums all the cubed values from 1 to n, and returns that sum.

Assume that the input n will always be a positive integer.

Examples:

SumCubes(2) == 9
// sum of the cubes of 1 and 2 is 1 + 8
*/

package kata

import (
	"math"
)

func SumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 3))
	}
	return sum
}
