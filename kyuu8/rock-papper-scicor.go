package kyuu8

import "fmt"

//Let's play! You have to return which player won! In case of a draw return Draw!.
//
//Examples(Input1, Input2 --> Output):
//
//"scissors", "paper" --> "Player 1 won!"
//"scissors", "rock" --> "Player 2 won!"
//"paper", "paper" --> "Draw!"

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	fmt.Println(m[p1])
	if m[p1] == p2 {
		return "Player 2 won!"
	}

	return "Player 1 won!"
}
