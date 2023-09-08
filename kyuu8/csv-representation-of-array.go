package kyuu8

import (
	"fmt"
	"strings"
)

//Create a function that returns the CSV representation of a two-dimensional numeric array.
//
//Example:
//
//input:
//[[ 0, 1, 2, 3, 4 ],
//[ 10,11,12,13,14 ],
//[ 20,21,22,23,24 ],
//[ 30,31,32,33,34 ]]
//
//output:
//'0,1,2,3,4\n'
//+'10,11,12,13,14\n'
//+'20,21,22,23,24\n'
//+'30,31,32,33,34'
//Array's length > 2.

//var _ = Describe("Tests", func() {
//	It("Sample tests", func() {
//		Expect(ToCsvText([][]int{
//			{0, 1, 2, 3, 45},
//			{10, 11, 12, 13, 14},
//			{20, 21, 22, 23, 24},
//			{30, 31, 32, 33, 34}})).To(Equal("0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34"))
//		Expect(ToCsvText([][]int{
//			{-25, 21, 2, -33, 48},
//			{30, 31, -32, 33, -34}})).To(Equal("-25,21,2,-33,48\n30,31,-32,33,-34"))
//		Expect(ToCsvText([][]int{
//			{5, 55, 5, 5, 55},
//			{6, 6, 66, 23, 24},
//			{666, 31, 66, 33, 7}})).To(Equal("5,55,5,5,55\n6,6,66,23,24\n666,31,66,33,7"))
//	})
//})

func ToCsvText(array [][]int) string {
	array2 := []string{}
	for _, a := range array {
		k := []string{}
		for _, n := range a {
			k = append(k, fmt.Sprint(n))
		}
		array2 = append(array2, strings.Join(k, ","))
	}
	return strings.Join(array2, "\n")
}
