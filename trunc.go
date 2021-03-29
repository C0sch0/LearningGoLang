package main

import (
	"fmt"
)

func main()  {
	var user_float float64

	fmt.Println("Enter float: ")
	fmt.Scan(&user_float)
	var user_int int = int(user_float)

	fmt.Printf("Int version is: %v ", user_int)
}