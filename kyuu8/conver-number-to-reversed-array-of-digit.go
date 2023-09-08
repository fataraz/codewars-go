package kyuu8

//Convert number to reversed array of digits
//Given a random non-negative number, you have to return the digits of this number within an array in reverse order.
//
//Example(Input => Output):
//
//35231 => [1,3,2,5,3]
//0 => [0]

func Digitize(n int) (arr []int) {
	for {
		arr = append(arr, n%10)
		n /= 10
		if n == 0 {
			return arr
		}
	}
}
