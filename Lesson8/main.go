package main

import (
	"fmt"
)

func main() {
	// Задача 8.1
	m1 := map[string]struct{}{"слон": {}, "бегемот": {}, "носорог": {}, "лев": {}}
	fmt.Println(m1)

	// Задача 8.2
	m2 := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}
	print2 := func(k string) {
		v, ok := m2[k]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", k, v, ok)
	}
	print2("слон")
	print2("бегемот")
	print2("осьминог")

	// Задача 8.3
	m3 := map[string]struct{}{"слон": {}, "бегемот": {}, "носорог": {}, "лев": {}}
	delete(m3, "бегемот")
	fmt.Println(m3)

	// Задача 8.4
	m4 := map[string]struct{}{"слон": {}, "бегемот": {}, "носорог": {}, "лев": {}}
	m4["выдра"] = struct{}{}
	fmt.Println(m4)

	// Задача 8.5
	m5 := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}
	m5["бегемот"] = 2
	fmt.Println(m5)

}
