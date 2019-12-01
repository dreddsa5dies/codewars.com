/*
This time no story, no theory. The examples below show you how to write function accum:

Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/

package kata

import "strings"

func Accum(s string) string {
	k := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[k] = r
		k++
	}
	rune = rune[0:k]

	final := ""
	for i := 1; i <= len(rune); i++ {
		tmp := strings.ToLower(string(rune[i-1]))
		final += strings.ToUpper(tmp) + strings.Repeat(tmp, i-1) + "-"
		if i == len(rune) {
			final = strings.TrimSuffix(final, "-")
		}
	}
	return final
}
