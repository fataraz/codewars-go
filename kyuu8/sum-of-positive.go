package kyuu8

//You get an array of numbers, return the sum of all of the positives ones.
//
//Example [1,-4,7,12] => 1 + 7 + 12 = 20
//
//Note: if there is nothing to sum, the sum is default to 0.

func PositiveSum(numbers []int) (result int) {
	for _, v := range numbers {
		if v > 0 {
			result += v
		}
	}
	return
}
