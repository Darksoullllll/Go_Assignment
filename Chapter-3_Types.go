package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Different Data Types

	//boolean Type
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
	println("=========================================================================")

	//Q2 Signed integer
	var a1 int = 89
	b1 := 95
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b1, unsafe.Sizeof(b1)) //type and size of b
	println()
	println("================================================================")
	//Floating Point Type
	a3, b3 := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	println("========================================================")
	//Complex Type
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	println("=================================================")
	//String Data Type
	first := "Abhinav"
	last := "Gautam"
	name := first + " " + last
	fmt.Println("My name is", name)
	println("==================================================")

	//Type Casting
	i := 55            //int
	j := 67.8          //float64
	sum1 := i + int(j) //j is converted to int
	fmt.Println(sum1)
}
