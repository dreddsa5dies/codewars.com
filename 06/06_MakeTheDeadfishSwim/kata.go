/*
Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/

package kata

func Parse(data string) []int {
	arr := []int{}
	val := 0
	for _, v := range data {
		switch {
		case string(v) == "i":
			val++
		case string(v) == "d":
			val--
		case string(v) == "s":
			val = val * val
		case string(v) == "o":
			arr = append(arr, val)
		default:
			continue
		}
	}
	return arr
}
