package main

// 1. Make program print how many tries it took to win. (try++)
// 2. See if you can tell if the user is lying.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//bufio is a package in core golang that can read from input sources.
	//we are saying build new one. we are saying input source is when people type.

	//1-100 is a our solution space, size of space we are searching - 100 items.

	// 100 ~ 100/2 guesses before it gets to answer because its starting in middle and guessing up or down by 1

	// n = 1 to 100 	0(n) -- linear is current approach (below)

	//binary search
	//search 1 -------- 50 -------- 100 - half it each time based on result
	//so if too high, then search between 1 and 50 (25)
	//if still too high, then between 1 and 25 (12.5)
	//this is optimum way to get answer
	// n -> binary search log(n)
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	var try int

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan() // It's going to wait until an input is received (by pressing Enter)

	// guess := 50 inefficient - dont use 'magic numbers'.
	for {
		try++
		guess := (low + high) / 2
		fmt.Println("This is try:", try)
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()             //Wait for response
		response := scanner.Text() //Take the text input

		//handle all possible responses with conditionals
		//now we are doing a binary search
		//the commented out was a 'linear' search

		if response == "a" {
			high = guess - 1
			if high == 0 {
				fmt.Println("You are cheating.")
				break
			}
			//guess--
		} else if response == "b" {
			low = guess + 1
			if low == 101 {
				fmt.Println("You are cheating.")
				break
			}
			//guess++
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invalid response, try again.")
		}

	}

}
