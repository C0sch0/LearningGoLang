package main

import (
	"fmt"
)
/*

 Module 2 Function Types
 Topic 1.1: First-class Value
-----------------------------------------------------

Way's that GoLang let's us use functions as first-class values
First class values are associated with functional programming.
There are some features of FP in GoLang which are very helpful.
Treating a function like a FCV is being able to treat functions like a variable of any other type

- Variables can be declared with a function type, just as you declare ints, strings, etc
- Can be created dynamically
- Can be passed as arguments and returned as values
- Can be stored in data structures

We will be able to implement this to function variables
- Declare a variable as a func

the variable starts being an alias for that function.

Example:
var funcVar func(int) int

(We use a function signature to indicate inputs and outputs)

func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar = incFn
	fmt.Print(funcVar(1))
}

Function is on right hand side, without().
This is indicating that funcVar is a pointer to that function.

----------------------------

Functions can be passed to another function as an argument

func applyIt(afunct func (int) int,
			val int) int {
	return afunct(val)
}

func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}

func main {
	fmt.Println(applyIt(incFn, 2)   // prints 3
	fmt.Println(applyIt(decFn, 2)	// prints 1
}


Anonymous Functions
-------------------------------------

Don't need to name a function. It's often useful to give them names to reference them.
(not naming comes from Lambda calculus)


-------------------------------------
example:

func applyIt(afunct func (int) int,
			val int) int {
	return afunct(val)
}

func main() {
	v := applyIt(func (x int) int {return x + 1}, 2)
	fmt.Println(v)
}
-------------------------------------


Functions as Return Values
Why ? If we want to make a new function with special parameters according to some data, parameterizable functions.

- Functions can return functions
- Might create a function with controllable parameters

Example:
- Distance to origin. Takes a point X,Y, Coordinates.
- Returns a distance to origin

What if I want to change the origin ?
Option 1: Pass origin as argument
Option 2: Define function with new origin

Function Defines a Function and returns a Function

func MakeDistOrigin (o_x, o_y float64) 													// Function returns functions
			func (float64, float64) float64 {											// Returning a function
	fn := func (x, y float64) float64 {													// Function we are creating
		return math.Sqrt(math.Pow(x - o_x, 2) + match.Pow(y - o_y, 2))					// Operation
	}
	return fn
}

- Origin location is passed as an argument
- Origin is built into the returned function


func main() {
	Dist1 := MakeDistOrigin(0,0)
	Dist2 := MakeDistOrigin(2,2)
	fmt.Println(Dist1(2,2))
	fmt.Println(Dist2(2,2))
}

--------------------------------------------


Environment of a Function: Set of all names that are valid inside a function
- Includes names defined locally, in the function
- Lexical scope: the variable can access variables defined in the block
- Environment includes names defined in block where function is defined

var x int
funct foo(y int) {
	z := 1
}




Closure
--------
- It is a function + its environment
- When functions are passed/returned, their environment comes with them !!
-




-------------------------------------------------------------
Module 2: Functions and Organization
Variadic and Deferred

Variable Argument Number
- Functions can take a variable number of arguments
- Use ellipsis ( ... ) to specify
- Treated as a slice inside a function


func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

You can treat the input as a slice


Variadic Slice Argument

func main() {
	fmt.Println(getMax(1, 3, 6, 4))
	vslice := []int{1, 3, 6, 4}
	fmt.Println(getMax(vslice...))
}

- Can pass a slice to a variadic function
- Need the ... suffix



Deferred Function Calls
They get executed later

- Call can be deferred until the surrounding function completes
- Typically used for cleanup activities

func main() {
	defer fmt.Println("Bye!")
	fmt.Println("Hello!")
}

We will see "Hello" and then "Bye"

Arguments of a deferred call are evaluated immediately, not later. Only the call is deferred, not the evaluation.

func main() {
	i := 1
	defer fmt.Println(i+1)
	i++
	fmt.Println("Hello!")
}

It prints a 2 because it's evaluated in that context, not later.


*/

func main() {
	fmt.Println("hi")
}