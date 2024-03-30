package main

import "fmt"

func main() {
	const (
		state1 = iota + 10
		state2
		state3
		state4
		state5
	)

	fmt.Println(state1, state2, state3, state4, state5)
}
