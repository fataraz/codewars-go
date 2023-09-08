package kyuu8

import "math"

//Complete the square sum function so that it squares each number passed into it and then sums the results together.
//
//For example, for [1, 2, 2] it should return 9 because
//1
//2
//+
//2
//2
//+
//2
//2
//=
//9
//1
//2
//+2
//2
//+2
//2
//=9.

func SquareSum(numbers []int) (result int) {
	for _, v := range numbers {
		result += int(math.Pow(float64(v), 2))
	}
	return
}
