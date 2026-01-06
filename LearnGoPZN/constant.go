package main

func main() {
	// const firstName string = "John"
	// const lastName string = "Doe"
	// const age int = 30

	// firstName = "Henry"
	// lastName = "Smith"

	// multi const
	const (
		firstName string = "John"
		lastName = "Doe"
		age = 20
	)

	println("First Name:", firstName)
	println("Last Name:", lastName)
	println("Age:", age)
	println("Full Name:", firstName+" "+lastName)
	println("Age in months:", age*12)
}
