package main

import (
	"fmt"
)

func main() {

	guessNumber()

}

func guessNumber() {

	fmt.Println("Pick a number between 0 - 100.")
	var ai int32 = 50
	var max int32 = 101
	var min int32 = 0
	var input string = ""
	var message string = ""

	for input != "yes" {

		message = fmt.Sprintf("Is your number %v ? Enter higher, lower, or yes\n", ai)
		fmt.Println(message)
		fmt.Scanln(&input)

		if input == "higher" {
			min = ai
			ai = (max + min) / 2
			fmt.Println(ai)
		} else if input == "lower" {
			max = ai
			ai = (max + min) / 2
		} else {
			fmt.Println("A.I. wins again")
		}
	}
}
