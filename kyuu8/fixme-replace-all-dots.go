package kyuu8

import "strings"

//The code provided is supposed replace all the dots . in the specified String str with dashes -
//
//But it's not working properly.
//
//Task
//Fix the bug so we can all go home early.
//
//Notes
//String str will never be null.

func ReplaceDots(str string) string {
	return strings.ReplaceAll(str, ".", "-")
}
