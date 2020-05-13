/*
If　a = 1, b = 2, c = 3 ... z = 26

Then l + o + v + e = 54

and f + r + i + e + n + d + s + h + i + p = 108

So friendship is twice stronger than love :-)

The input will always be in lowercase and never be empty.
*/

package kata

func WordsToMarks(s string) int {
	r := []rune(s)
	c := 0
	for _, v := range r {
		c += int(v) - 96
	}
	return c
}
