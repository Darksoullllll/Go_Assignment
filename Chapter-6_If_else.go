package main

import "fmt"

func main() {
	//If_Else
	println("Odd Even number")
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")

	println("===============================")

	println("Maximum of two number")
	num1 := 30
	num2 := 78
	if num1 > num2 {
		println("The Maximum Number is: ", num1)
	} else {
		println("The Maximum Number is: ", num2)
	}
	println("======================================")
	n := 99
	if n <= 50 {
		fmt.Println(n, "is less than or equal to 50")
	} else if n >= 51 && n <= 100 {
		fmt.Println(n, "is between 51 and 100")
	} else {
		fmt.Println(n, "is greater than 100")
	}
}
