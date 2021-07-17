package main

import (
	"fmt"
)

func foo(x *[3]int) int {
	(*x)[0] = (*x)[0] * 1
	return 0
}

func main(){
	a := [3]int{1, 2, 3}
	foo(&a)
	fmt.Print(a)
}


// messy and unnecessary
// the way to do it in go is using SLICES. Get used to use slices instead of arrays.

// Slice contains a pointer to the array
// Passing a slice copies the pointer
// Slice is like a window on an array

// if you describe a slice from scratch, it will make an array behind scenes
// A slice is a structure that contains a pointer to the array or to the start of the array, the lenght and the capacity
// it's still called by value but we are copying the pointer.
//

func foo2(sli int) int {
	sli[0] = sli[0] + 1
}

func main2()  {
	a := []int{1, 2, 3}
	foo2(a)
	fmt.Print(a)
}

// when we print a, it will print out the changed slice
// when you want to pass an array arg, pass a slice better so we don't have to do the referencing ourselves.
