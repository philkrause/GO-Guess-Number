package main

import "fmt"

func main() {

	var ai int32 = 50
	var max int32 = 100
	var min int32 = 0
	var input string = ""
	var message = fmt.Sprintf("Is your number %d Enter higher, lower, or yes\n", ai)

	fmt.Println("The Number Guesser in GO!")
	fmt.Println(message)

	for input != "yes" {

		fmt.Scanln(&input)
		fmt.Println(message)

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
