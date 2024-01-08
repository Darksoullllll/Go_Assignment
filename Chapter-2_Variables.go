package main

import (
	"fmt"
	"math"
)

func main() {
	// declaring a variable

	var age int
	println("My age is : ", age)
	println("=====================")

	//Q2
	// declaring a new variable
	var age2 int
	println("My age is: ", age2)
	age2 = 20 // assigment
	println("My age is: ", age2)
	age2 = 25
	println("My new age is: ", age2)

	println("=========================")

	//Q3
	var width, height int = 100, 50 //declaring multiple variables

	println("width is", width, "height is", height)
	println("===================================================")

	var (
		name     = "Abhinav Gautam"
		myage    = 20
		myheight int
	)
	println("my name is", name, ", age is", myage, "and height is", myheight)
	println("=================================================================")

	a, b := 20, 30 // declare variables a and b
	println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	println("changed b is", b, "c is", c)

	println("======================================================")
	a1, b1 := 145.8, 543.8
	c1 := math.Min(a1, b1)
	fmt.Println("Minimum value is", c1)
}
