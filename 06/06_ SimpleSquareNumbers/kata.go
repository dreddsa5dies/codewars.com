/*
In this Kata, you will be given a number n (n > 0) and your task will be to return
the smallest square number N (N > 0) such that n + N is also a perfect square. If
there is no answer, return -1 (nil in Clojure, Nothing in Haskell).
*/

package kata

func Solve(n int) int {
	fin := -1

	for i := n*n - 1; i > 0; i-- {
		if n%i == 0 {
			bma := n/i - i
			if bma%2 == 0 {
				p := bma / 2
				if p > 0 {
					fin = p * p
					break
				}
			}
		}
	}

	return fin
}
