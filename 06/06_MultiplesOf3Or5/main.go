package kata

func Multiple3And5(number int) int {
	if number <= 0 {
		return 0
	}
	fin := 0

	for i := 1; i < number; i++ {
		switch {
		case i%3 == 0 || i%5 == 0:
			fin += i
		}
	}
	return fin
}
