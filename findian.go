package main

// Write a program which prompts the user to enter a string.
//The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
//The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’,
//and contains the character ‘a’. The program should print “Not Found!” otherwise.
//The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

import (
	"fmt"
)


func check_for_a (user_input string) bool {
	for _, char := range user_input {
		if char == 'a' {
			return true
		}
	}
	return false
}


func main()  {
	var user_string_input string
	fmt.Println("Enter string: ")
	fmt.Scan(&user_string_input)

	if user_string_input[len(user_string_input)-1:] != "n" {
		fmt.Println("Not Found!")
		return
	} else if user_string_input[0:1] != "i" {
		fmt.Println("Not Found!")
		return
	} else if !check_for_a(user_string_input) {
		fmt.Println("Not Found!")
		return
	}
	fmt.Println("Found!")
}

