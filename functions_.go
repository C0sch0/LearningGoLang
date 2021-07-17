package main


/*
Everything here
will be considered
a block comment

//
// Understandability
// We need a well organized code
// It's important to understand our own code, for debugging, maintenance
// the way we write our functions really impacts Understandability

// Code in functions and data
// things we're gonna change and how we are going to change them

// this program is understandable

//finding feature quickly,
// If we order information is easier to find things
// we need to be able to find things easily

// We want to be able to trace through data

// Which piece of code is affecting our information

// Debugging principles

// Code crashes inside a function.
Two options for the cause

1. Function is written incorrectly
Sorts a slice in the wrong order

2. Data that the function uses is incorrect.
Sorts slice correctly but slice has wrong elements in it

Inputs are giving from parameters, files, user input, etc

Supporting Debugging

Functions need to be understandable, is this right, wrong ?
Does it behave as we need

Data needs to be traceable
- Where did that data come from ?
Global variables complicate this. They could come from anywhere, so it's hard to trace.
We still need them tho

Module 1, Topic 2.2 Guidelines for functions

Making good functions.
Debugging, collaboration, upgrading, etc

1. Function Naming
Give functions a good name
Behaviour can be understood at a glance
Parameter naming counts too

func ProcessArray (a []int) float {}

versus

func ComputeRMS (samples []float) float {}

Name gives information about what's going to happen - this is what we need

Don't make it too long tho !!

Correct variable and function naming makes everyone happy :)

Functional Cohesion

Functions should perform only one operation

The process depends on the context.

Example: Geometry application

Good functions
	PointDist(), DrawCircle(), TriangleArea

Bad functions: Merging behaviors makes code complicated
 - DrawCircle() + TriangleArea


Few parameters
Debugging requires tracing function input data

More difficult with a large number of parameters


Functions should be simple, easier to debug.
- Short functions can be complicated too


How do you write complicated code with simple functions ?
Function Call Hierarchy

Option 1, one big func, 100 lines
Op. 2, one functin that calls two 50 line functions.


Control-flow Complexity

Control flow describes conditional paths.

- 3 control-flow

Functional hierarchy can reduce control flow complexity



*/