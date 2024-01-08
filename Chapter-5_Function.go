package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	price, no := 150, 5
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(5.6, 2.5)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	println()
	area2, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area2)
}
