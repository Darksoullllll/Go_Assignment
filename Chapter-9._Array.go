package main

import "fmt"

func main() {
	var a [3]int //int array with length 3
	fmt.Println(a)

	var a1 [3]int //int array with length 3
	a1[0] = 12    // array index starts at 0
	a1[1] = 78
	a1[2] = 50
	fmt.Println(a1)

	arr := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(arr)

	ar := [3]int{12}
	fmt.Println(ar)

	a2 := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a2)

	a3 := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a3 // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a3)
	fmt.Println("b is ", b)

	a4 := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a4[i])
	}

	a5 := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a5 { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	println()
	println("Array Slicing")
	a6 := [5]int{76, 77, 78, 79, 80}
	var b1 []int = a6[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b1)

	println()
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	println()
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	println()
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
}
