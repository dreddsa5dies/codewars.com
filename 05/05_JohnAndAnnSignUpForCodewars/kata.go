package kata

func Ann(n int) []int {
	tmp := []int{}
	if n == 0 {
		return []int{1}
	}
	a := 1
	for i := 0; i < n; i += 2 {
		tmp = append(tmp, a, a)
		a++
	}
	return tmp
}

func John(n int) []int {
	tmp := []int{}
	for i := 0; i <= n; i++ {
		if i < 2 {
			tmp = append(tmp, 0)
		} else {
			ann := Ann(tmp[i-2])
			tmp = append(tmp, i-ann[len(ann)-1])
		}
	}
	return tmp
}

func SumJohn(n int) int {
	tmp := John(n)
	count := 0
	for _, v := range tmp {
		count += v
	}
	return count
}

func SumAnn(n int) int {
	tmp := Ann(n)
	count := 0
	for _, v := range tmp {
		count += v
	}
	return count
}
