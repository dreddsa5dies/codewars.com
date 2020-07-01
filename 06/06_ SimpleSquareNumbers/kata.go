/*
In this Kata, you will be given a number n (n > 0) and your task will be to return
the smallest square number N (N > 0) such that n + N is also a perfect square. If
there is no answer, return -1 (nil in Clojure, Nothing in Haskell).
*/

package kata

func Solve(n int) int {
	res := -1
	for i := 1; i*i < n; i++ {
		if n%i == 0 && (n/i-i)%2 == 0 {
			res = (n/i - i) * (n/i - i) / 4
		}
	}
	return res
}
