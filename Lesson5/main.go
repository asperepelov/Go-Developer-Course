package main

import (
	"fmt"
	"strconv"
)

type square int

func main() {
	//Задача 5.1
	var ptrStr *string
	fmt.Println(ptrStr)

	//Задача 5.2
	v2 := "good"
	fmt.Println(v2, &v2)

	//Задача 5.3
	v3 := "good"
	ptr := &v3
	*ptr = "bad"
	fmt.Println(v3)

	//Задача 5.4
	var v41, v42, v43, v44 int16 = 1, 2, 3, 4
	fmt.Println(v41, v42, v43, v44)
	fmt.Println(&v41, &v42, &v43, &v44)

	//Задача 5.5
	var v5 int = 10
	change(&v5)
	fmt.Println(v5)

	//Задача 5.6
	var s6 square = 25
	fmt.Println(s6)

	//Задача 5.7
	var s7 square = 30
	s7 += 15
	fmt.Println(s7)

	//Задача 5.8
	var s8 square = 34
	s8 += 10
	fmt.Println(toStr(s8))

}

func change(p *int) {
	*p++
}

func toStr(s square) string {
	return strconv.Itoa(int(s)) + " м²"
}
