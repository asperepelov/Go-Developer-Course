package main

import "fmt"

const (
	state1 = iota + 1
	state2
	state3
	state4
	state5
)

func main() {
	fmt.Println(state1, state2, state3, state4, state5)
}
