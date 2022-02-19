package main

import (
	"fmt"
	"time"
)

var test = "blah"

func getData(customerId int) (customer string) {
	if customerId == 1 {
		return "Andy Godish"
	} else if customerId == 2 {
		return "Bob Smith"
	} else {
		return ""
	}
}

func getCustomers() (customers [2]string) {
	customer := "Andy Godish"
	c2 := "John"

	customers[0] = customer
	customers[1] = c2

	return customers
}

func getSlice() (customers []string) {
	customers = []string{"test user"}

	customers = append(customers, "Andy Godish", "John", "Fred", "Heyo")

	for {
		fmt.Println("Infinite Loop 1")
		time.Sleep(time.Second)
		break
	}

	for x := 0; x < len(customers); x++ {
		fmt.Println("Loop 2" + " " + customers[x])
		time.Sleep(time.Second)
		if x == 3 {
			break
		}
	}

	for _, customer := range customers {
		fmt.Println("loop 3" + " " + customer)
	}

	return customers
}

func main() {
	fmt.Println(getData(2))
	fmt.Println(getData(1))
	fmt.Println(test)

	fmt.Println(getCustomers())
	fmt.Println(len(getCustomers()))

	fmt.Println(getSlice())

	customers := GetAllCustomers()
	fmt.Println(customers)
}
