package kyuu8

//Given a non-empty array of integers, return the result of multiplying the values together in order. Example:
//[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

func Grow(arr []int) (result int) {
	result = 1

	for _, n := range arr {
		result *= n
	}

	return
}
