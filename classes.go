package main

import (
	"fmt"
)

/*
Classes
A class is a collection of data fields and functions that share a well-defined responsibility

Data + functions (methods) put together.

Example:
Point class: Used in a geometry program
- Data: X coordinate, y coordinate
Functions: DistToOrigin, Quadrant, AddXOffset, AddYOffset, SetX, SetY

Classes are a template
Contain data fields, not data


An object is an instance of a class

They use the template and contain real data

Class != Object

Encapsulation
- Data can be protected from the programmer

We might want to hide some data
- Data can be accessed only using methods

- Maybe we don't trust the programmer to keep data consistent. We relieve the programmer from the burden of consistency.
Through these methods, we ensure consistency is kept.

Example: Double distance to origin
Option 1: Make method DoubleDist
Option 2: Trust programmer to double X and Y directly: This has the risk of potential mistakes


------------------------------------------
Module 3
------------------------------------------
Object-Orientation in Go
Topic 1.2: Support for Classes
------------------------------------------

GoLang has no "Class" Keyword.
Most OO languages have a class keyword

In Go, we have a different way to associate data with methods.

Associating methods with Data is done with receiver types.
Receiver type is the type the method is associated with.

Method has a receiver type that it is associated with
Use dot notation to call the method


As a rule: Whenever you associate a method with a type, you need them to be in the same packages.
*/

type MyInt int

func (mi MyInt) Double () int {
	return int(mi*2)
}
// Double only works with MyInt. MyInt is the receiver type, and mi references the specific MyInt we will work with
// Receiver type is defined before the function name.

func main()  {
	v := MyInt(3)
	fmt.Println(v.Double())
}
// V is passed implicitly as an argument when it's called like this. In this case, Double has 1 argument.
// pass by value !! A Copy of it is passed
//



/*


Implicit Method Argument
- Object v is an implicit argument to the method
- - Call by value


*/


/*

------------------------------------------
Module 3
------------------------------------------
Object-Orientation in Go
Topic 1.3: Support for Classes
------------------------------------------

Usually you can put a lot of data and data types and associate with methods.
In Go we can do the same, but instead of classes we use receiver types.

A way to associate multiple types of data with functions, we use Structs as receiver types

Structs types compose data fields

type Point struct {
	x float64
	y float64
}

Traditional future: using Structs as receivers


Structs with Methods
Structs and methods together allow arbitrary data and functions to be composed

func (p Point) DistToOrig() {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	p1 := Point(3, 4)
	fmt.Println(p1.DistToOrig())           // implicitly passing p1 to DistToOrig
}

*/


/*

------------------------------------------
Module 3
------------------------------------------
Object-Orientation in Go
Topic 2.3: Encapsulation
------------------------------------------

Controlling Access
- Can define public functions to allow access to hidden data

package data
var x int= 1
func PrintX() {fmt.Println(x)}

PrintX starts with Capital letter, which means it gets exported, everyone can access this function
If we import the data package, we will be able to access that function, even though we can't access x

package main
import "data"

func main() {
	data.PrintX()
}


We let people to see X but only through the access we give them, without being able to change it

We can do this with structures too.

Controlling Access to Structs

Hide fields of structs by starting field name with a lower-case letter
yet, we define public methods which access hidden data

package data
type Point struct{
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float 64){
	p.x = xn
	p.y = yn
}

This way, we let other devs see our data in a controlled way

*/


/*

------------------------------------------
Module 3
------------------------------------------
Object-Orientation in Go
Topic 2.3: Point Receivers
------------------------------------------

Limitations of Methods
- Receiver is passed implicitly as an argument to the method
- Method cannot modify data inside the receiver

Example: OffsetX() should increase x coordinate

func main() {
	p1 := Point(3, 4)
	p1.OffsetX(5)
}

- How do we change the actual values in p1 ? We can't. Because we are passing a copy, not the original.
- Large Receivers: when we pass large objects we face problems of speed and memory

How do we avoid this ?
We pass by reference, not by value.

Pointer Receivers
-------------------------

func (p *Point) OffsetX(v float64)
{
	p.x = p.x + v
}

Receiver can be a pointer to a type.
Call by reference, pointer is passed to the method. Now we can modify data.



*/


/*

------------------------------------------
Module 3
------------------------------------------
Object-Orientation in Go
Topic 2.3: Point Receivers, Referencing and Dereferencing
------------------------------------------

One thing with using a pointer is that there is no need to dereference the pointer

func (p *Point) OffsetX(v int){
	p.x = p.x + v
}

Point is referenced as p, not *p, right ?
Dereference is automatic with . operator. GoLang knows that you want this.

Also, no need to reference.

func main() {
	p := Point{3, 4}
	p.OffsetX(5)
	fmt.Println(p.x)
}

p is a struct. When we call OffsetX, we give p.OffsetX(5), we say "p.", so you would thing you need &p. But no. Cool.

When using pointer receivers, it is a good programming practice
1. All methods for a type have pointer receivers, or
2. All methods for a type have non-pointer receivers

Mixing pointer/non-pointers for a type will get confusing

*/