package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
	Go does not have traditional inheritance like some
	object-oriented languages where one class inherits fields
	and methods from another. Instead, Go uses composition, which
	allows you to embed one struct inside another.
*/

type Employee struct {
	Person
	Position string
}

type Speak interface {
	SayHello() string
}

func (p Person) SayHello() string {
	return "Hi, my name is " + p.Name
}

func saySomething(s Speak) {
	fmt.Println(s.SayHello())
}

func main() {
	p := Person{Name: "Henry", Age: 20}
	fmt.Printf("Name: %s, age: %d\n", p.Name, p.Age)

	e := Employee{
		Person:   Person{Name: "Salim", Age: 35},
		Position: "Manager",
	}

	// bisa tanpa .Person juga kalau mau access name sama age
	fmt.Printf("Name: %s, age: %d, position: %s\n", e.Person.Name, e.Person.Age, e.Position)

	saySomething(Person{})
	p = Person{Name: "Alice", Age: 30}
	saySomething(p)
}
