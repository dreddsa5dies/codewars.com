/*
Your task is to add up letters to one letter.

The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

Notes:
Letters will always be lowercase.
Letters can overflow (see second to last example of the description)
If no letters are given, the function should return 'z'
Examples:
AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'
Confused? Roll your mouse/tap over here
*/

package kata

func AddLetters(letters []rune) rune {
	alphabetMap := make(map[rune]int)
	alphabet := [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for k, v := range alphabet {
		alphabetMap[v] = k + 1
	}

	var final rune

	switch {
	case len(letters) == 0:
		final = 'z'
	case len(letters) == 1:
		final = letters[0]
	default:
		count := 0
		for i := 0; i <= len(letters)-1; i++ {
			count += alphabetMap[letters[i]]
		}
		count = count % 26
		if count == 0 {
			final = 'z'
		}
		for k, v := range alphabetMap {
			if v == count {
				final = k
			}
		}
	}
	return final
}
