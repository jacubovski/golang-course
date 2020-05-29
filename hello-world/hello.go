package main

import "fmt"

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("something more")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("world")
}

func bar() {
	fmt.Println("and then we exited")
}
