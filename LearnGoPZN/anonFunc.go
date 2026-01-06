package main

import "fmt"

func a() func() {
    n := 0
    return func() {
        n++
        fmt.Println(n)
    }
}

func main() {
	func() {
		fmt.Println("This is anon function")
	}()

	f := func ()  {
		fmt.Println("This is anon function inside variable")
	}

	g := func(v int) {
		fmt.Println(v)
	}

	f()
	g(42)

	func(n int) {
		fmt.Println(n)
	}(52)

	func(v int, g func(v int)) {
		g(v)
	}(5, g)

	callA := a()
	callA()
}
