package main

import "fmt"

func zeroValue() {
	// a = 0, s = ""
	var a, b int
	var s string
	fmt.Println("1. ", a, b, s)
	fmt.Printf("2. a = %d, b = %d, s = %s\n", a, b, s)
	// %q print empty string as ""
	fmt.Printf("3. a = %d, b = %d, s = %q\n", a, b, s)
}

func initialValue() {
	var a, b int = 1, 2
	var s string = "hello"

	fmt.Println(a, b, s)
}

func auto() {
	var a, b, bool_, s = 1, 2, true, "hello"
	var (
		c int = 3
		s2 string = "world"
	)
	fmt.Println(a, b, bool_,s)
	fmt.Println(c, s2)
}

func shortcut() {
	a, b, bool_, s := 1, 2, true, "hello"
	fmt.Println(a, b, bool_,s)
}

func main() {
	zeroValue()
	fmt.Println("\ninitial values: ")
	initialValue()
	fmt.Println("\nauto: ")
	auto()
	fmt.Println("\nshortcut:")
	shortcut()
}
