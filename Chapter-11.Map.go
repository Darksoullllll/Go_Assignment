package main

import "fmt"

func main() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)

	//Adding element in map
	employeeSalary["Abhinav"] = 12000
	employeeSalary["xyz"] = 15000
	employeeSalary["abc"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	println()
	employeeSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary2["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary2)

	println()
	/*var employeeSalary3 map[string]int
	employeeSalary3["steve"] = 12000*/
	println("Retrieving value for a key from a map")
	employee := "Abhinav"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	println("this print 0")
	fmt.Println("Salary of joe is", employeeSalary["joe"])

	println()
	println("Checking if a key exists")
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	println()
	println("Iterate over all elements in a map")
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	println()
	println("Deleting items from a map")
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
}
