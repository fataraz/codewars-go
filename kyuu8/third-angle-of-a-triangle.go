package kyuu8

//You are given two interior angles (in degrees) of a triangle.
//
//Write a function to return the 3rd.
//
//Note: only positive integers will be tested.
//
//https://en.wikipedia.org/wiki/Triangle

//var _ = Describe("Basic Tests", func() {
//	It("OtherAngle(30, 60)", func() { Expect(OtherAngle(30, 60)).To(Equal(90)) })
//	It("OtherAngle(60, 60)", func() { Expect(OtherAngle(60, 60)).To(Equal(60)) })
//	It("OtherAngle(43, 78)", func() { Expect(OtherAngle(43, 78)).To(Equal(59)) })
//	It("OtherAngle(10, 20)", func() { Expect(OtherAngle(10, 20)).To(Equal(150)) })
//})

func OtherAngle(a int, b int) int {
	return 180 - a - b
}
