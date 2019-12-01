/*
Let us begin with an example:

Take a number: 56789. Rotate left, you get 67895.

Keep the first digit in place and rotate left the other digits: 68957.

Keep the first two digits in place and rotate the other ones: 68579.

Keep the first three digits and rotate left the rest: 68597. Now it is over since keeping the first four it remains only one digit which rotated is itself.

You have the following sequence of numbers:

56789 -> 67895 -> 68957 -> 68579 -> 68597

and you must return the greatest: 68957.

Task
Write function max_rot(n) which given a positive integer n returns the maximum number you got doing rotations similar to the above example.

So max_rot (or maxRot or ... depending on the language) is such as:

max_rot(56789) should return 68957

max_rot(38458215) should return 85821534
*/

package kata

import "strconv"

func MaxRot(n int64) int64 {
	tmpStr := strconv.FormatInt(n, 10)
	k := 0
	rune := make([]rune, len(tmpStr))
	for _, r := range tmpStr {
		rune[k] = r
		k++
	}
	rune = rune[0:k]

	greatest := make([]int64, len(rune))

	for i := 0; i < len(rune); i++ {
		tmpRune := rune[i]
		rune = append(rune[:i], rune[i+1:]...)
		rune = append(rune, tmpRune)
		tmpInt, _ := strconv.Atoi(string(rune[0:len(rune)]))
		greatest[i] = int64(tmpInt)
	}
	greatest = append(greatest, n)

	var tmpInt64 int64
	for _, v := range greatest {
		if v > tmpInt64 {
			tmpInt64 = v
		}
	}

	return tmpInt64
}
