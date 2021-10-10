package mymodule

import "fmt"

func hello() string {
	return "Hello there!"
}

func add(addend1 int, addend2 int) int {
	return addend1 + addend2
}

func greet(name string) string {
	greeting := fmt.Sprintf("Hello %v, nice to see you!", name)
	return greeting
}

func iseven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
