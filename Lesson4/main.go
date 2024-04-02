package main

import "fmt"

func main() {
	//Задача 4.1
	hello()

	//Задача 4.2
	func() {
		fmt.Println("Hello Go!")
	}()

	//Задача 4.3
	f3 := func() {
		fmt.Println("Hello Go!")
	}
	hello3(f3)

	//Задача 4.4
	f4 := hello4()
	f4()

	//Задача 4.5
	hello5()

	//Задача 4.6 с использованием замыкания и рекурсии
	priorVal := 0
	curVal := 1
	cntVal := 23
	var fibonachi func()
	fibonachi = func() {
		if cntVal <= 0 {
			return
		}
		nextVal := priorVal + curVal
		priorVal, curVal = curVal, nextVal
		fmt.Println(nextVal)
		cntVal--
		fibonachi()
	}

	fibonachi()

}

func hello() {
	fmt.Println("Hello Go!")
}

func hello3(f func()) {
	f()
}

func hello4() func() {
	return func() {
		fmt.Println("Hello Go!")
	}
}

func hello5() {
	fmt.Println("Hello Go!")
	defer fmt.Println("Завершение hello5")
}
