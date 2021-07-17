package main

import (
	"fmt"
)


/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

func Swap(slice []int, pos int){
	tmpVar := slice[pos]
	slice[pos] = slice[pos + 1]
	slice[pos + 1] = tmpVar
}

func BubbleSort(slice []int) {
	swapNeeded := true
	for swapNeeded {
		swapNeeded = false
		for index := 0; index < len(slice) - 1; index++ {
			if slice[index] > slice[index + 1] {
				Swap(slice, index)
				swapNeeded = true
			}
		}
	}
}

func promptUser(intSlice []int)  {
	for i := 0; i < 10 ; i++ {
		var userInt int
		fmt.Printf("Please provide the number for index %d:  ", i)
		fmt.Scan(&userInt)
		intSlice[i] = userInt
	}
}

func main()  {
	intSlice := make([]int, 0, 10)
	// Prompt user for 10 integers
	promptUser(intSlice)
	// BubbleSort
	BubbleSort(intSlice)
	// Print those numbers, in order from least to greatest.
	fmt.Print(intSlice)
}
