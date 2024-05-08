package main

import "fmt"

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10
	if vv, ok := v.(int); ok {
		a += vv
	} else {
		fmt.Println("Не возможно v привести к int")
		return
	}

	fmt.Println(a)
}
