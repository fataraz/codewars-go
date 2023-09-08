package kyuu8

import (
	"strings"
)

//Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
//
//The output should be two capital letters with a dot separating them.
//
//It should look like this:
//
//Sam Harris => S.H
//
//patrick feeney => P.F
//
//var _ = Describe("Test Example", func() {
//	It("should test that the solution returns the correct value", func() {
//		Expect(AbbrevName("Sam Harris")).To(Equal("S.H"))
//		Expect(AbbrevName("Patrick Feenan")).To(Equal("P.F"))
//		Expect(AbbrevName("Evan Cole")).To(Equal("E.C"))
//		Expect(AbbrevName("P Favuzzi")).To(Equal("P.F"))
//		Expect(AbbrevName("David Mendieta")).To(Equal("D.M"))
//	})
//})

func AbbrevName(name string) string {
	words := strings.Split(name, " ")
	return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}
