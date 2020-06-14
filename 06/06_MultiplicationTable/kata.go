/*
Your task, is to create NxN multiplication table, of size provided in parameter.

for example, when given size is 3:

1 2 3
2 4 6
3 6 9
for given example, the return value should be: [[1,2,3],[2,4,6],[3,6,9]]
*/

package multiplicationtable

func MultiplicationTable(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}
	if size == 1 {
		return [][]int{{1}}
	}

	fin := [][]int{}
	for i := 1; i <= size; i++ {
		t := []int{}
		for k := 1; k <= size; k++ {
			t = append(t, i*k)
		}
		fin = append(fin, t)
	}

	return fin
}
