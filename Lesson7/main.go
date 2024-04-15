package main

import "fmt"

func main() {
	// Задача 7.1
	s1 := [5]string{"один", "два", "три", "четыре", "пять"}
	fmt.Println("s1 =", s1)

	// Задача 7.2
	s2 := [...]string{"яблоко", "груша", "слива", "абрикос", "фрукты"}
	fmt.Println("s2 =", s2)

	// Задача 7.3
	s3 := [...]string{"яблоко", "груша", "помидор", "абрикос"}
	s3[2] = "персик"
	fmt.Println("s3 =", s3)

	// Задача 7.4
	s4 := []int{5, 2, 8, 3, 1, 9}
	fmt.Println("s4 =", s4)

	// Задача 7.5
	s5 := make([]int, 0, 10)
	fmt.Println("s5 =", s5)

	// Задача 7.6
	s6 := make([]int, 0, 10)
	s6 = append(s6, 4, 1, 8, 9)
	fmt.Println("s6 =", s6)

	// Задача 7.7
	s7 := []int{1, 2, 3}
	s71 := []int{4, 5, 6}
	s7 = append(s7, s71...)
	fmt.Println("s7 =", s7)

	// Задача 7.8
	s8 := []int{1, 2, 3, 4, 5, 6}
	s8 = append(s8[:3], s8[4:]...)
	fmt.Println("s8 =", s8, "len(s8) =", len(s8), "cap(s8) =", cap(s8))

}
