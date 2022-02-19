package main

type Customer struct {
	FirstName string
	LastName  string
	Fullname  string
}

func GetAllCustomers() (customers []Customer) {
	customer := Customer{FirstName: "Andy", LastName: "Godish"}

	customers = append(customers, customer,
		Customer{FirstName: "Test", LastName: "User"},
	)

	return customers
}
