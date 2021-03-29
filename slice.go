package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	var int_slice = make([]int, 0, 3)

	for {
		var user_string_input string
		fmt.Println("Enter int to append or X to quit: ")
		fmt.Scan(&user_string_input)

		if user_string_input == "X"{
			break
		} else {
			to_int, _ := strconv.Atoi(user_string_input)
			int_slice = append(int_slice, to_int)
			sort.Ints(int_slice)
			fmt.Println(int_slice)
		}
	}
}
