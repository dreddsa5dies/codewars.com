/*
Do you know how to make a spiral? Let's test it!
Classic definition: A spiral is a curve which emanates from a central point, getting progressively farther away as it revolves around the point.

Your objective is to complete a function createSpiral(N) that receives an integer N and returns an NxN two-dimensional array with numbers 1 through NxN represented as a clockwise spiral.

Return an empty array if N < 1 or N is not int / number

Examples:

N = 3 Output: [[1,2,3],[8,9,4],[7,6,5]]

1    2    3
8    9    4
7    6    5
N = 4 Output: [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]

1   2   3   4
12  13  14  5
11  16  15  6
10  9   8   7
N = 5 Output: [[1,2,3,4,5],[16,17,18,19,6],[15,24,25,20,7],[14,23,22,21,8],[13,12,11,10,9]]

1   2   3   4   5
16  17  18  19  6
15  24  25  20  7
14  23  22  21  8
13  12  11  10  9
*/

package kata

func CreateSpiral(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	var col, row, nextCol, nextRow int
	clockwise := true
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	direction := nextDirection("", clockwise)
	sequenceIndex := 1
	result[row][col] = 1
	for {
		if sequenceIndex == n*n {
			break
		}
		nextCol, nextRow = next(direction, col, row)
		inside := isInBoundsAndNotVisited(result, nextCol, nextRow)
		if inside {
			col = nextCol
			row = nextRow
			sequenceIndex++
			result[row][col] = sequenceIndex
		} else {
			direction = nextDirection(direction, clockwise)
		}
	}
	return result
}

func isInBoundsAndNotVisited(array [][]int, col, row int) bool {
	if row >= 0 && row < len(array) {
		// fmt.Print("rowOk ", row)
		if col >= 0 && col < len(array) {
			return array[row][col] == 0
		}
	}
	return false
}

func nextDirection(direction string, clockwise bool) string {
	if clockwise {
		return nextClockwiseDirection(direction)
	}
	return nextCounterClockwiseDirection(direction)
}

func nextClockwiseDirection(direction string) string {
	switch direction {
	case "r":
		return "d"
	case "d":
		return "l"
	case "l":
		return "u"
	case "u":
		return "r"
	default:
		return "r"
	}
}

func nextCounterClockwiseDirection(direction string) string {
	switch direction {
	case "r":
		return "u"
	case "d":
		return "r"
	case "l":
		return "d"
	case "u":
		return "l"
	default:
		return "d"
	}
}

func next(direction string, col, row int) (newCol, newRow int) {
	newCol, newRow = col, row
	switch direction {
	case "r":
		newCol = col + 1
	case "d":
		newRow = row + 1
	case "l":
		newCol = col - 1
	case "u":
		newRow = row - 1
	}
	return
}
