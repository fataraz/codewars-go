package kyuu8

import "math"

//Your task is to find the nearest square number, nearest_sq(n) or nearestSq(n), of a positive integer n.
//
//For example, if n = 111, then nearest\_sq(n) (nearestSq(n)) equals 121, since 111 is closer to 121, the square of 11, than 100, the square of 10.
//
//If the n is already the perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.
//
//Good luck :)
//
//Check my other katas:
//
//Alphabetically ordered
//
//Case-sensitive!
//
//Not prime numbers
//
//Find your caterer
//
//var _ = Describe("Test", func() {
//	It("should test basic cases", func() {
//		Expect(NearestSq(1)).To(Equal(1))
//		Expect(NearestSq(2)).To(Equal(1))
//		Expect(NearestSq(10)).To(Equal(9))
//		Expect(NearestSq(111)).To(Equal(121))
//		Expect(NearestSq(9999)).To(Equal(10000))
//	})
//})

func NearestSq(n int) int {
	n = int(math.Round(math.Sqrt(float64(n))))

	return n * n
}
