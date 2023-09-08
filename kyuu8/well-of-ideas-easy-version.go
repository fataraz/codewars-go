package kyuu8

//For every good kata idea there seem to be quite a few bad ones!
//
//In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas,
//return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.

//var _ = Describe("Test Example", func() {
//	It("should test example values", func() {
//		Expect(Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "good", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "good", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "bad", "good", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "good"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "good", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "good", "good", "good", "bad", "bad", "good"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "good", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good", "bad", "bad", "bad", "bad", "good", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good", "good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "good", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good", "bad", "good", "good", "bad", "bad", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "good", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "good", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"})).To(Equal("Publish!"))
//		Expect(Well([]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad"})).To(Equal("Fail!"))
//		Expect(Well([]string{"good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "bad", "bad"})).To(Equal("I smell a series!"))
//		Expect(Well([]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "bad"})).To(Equal("Publish!"))
//	})
//})

func Well(x []string) string {
	var good int
	for _, v := range x {
		if v == "good" {
			good++
		}
	}
	if good == 0 {
		return "Fail!"
	}
	if good > 2 {
		return "I smell a series!"
	}

	return "Publish!"
}
