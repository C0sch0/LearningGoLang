package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	slice := make([]string, 3)
	slice[2] = "2"
	fmt.Println("emp:", slice)

	mapping := make(map[string]int)
	mapping["k1"] = 7
	mapping["k2"] = 13

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
