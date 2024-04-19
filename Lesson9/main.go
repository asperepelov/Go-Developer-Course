package main

import "fmt"

type fruit string
type count int

// Задача 9.1
func fruitMarket(f fruit) {
	fruits := map[fruit]count{"апельсин": 5, "яблоки": 3, "сливы": 1, "груши": 0}
	cnt, ok := fruits[f]
	if !ok {
		fmt.Printf("%s отсутствует в номенклатуре магазина!\n", f)
		return
	}

	if cnt == 0 {
		fmt.Printf("%s - %d, привозить больше?\n", f, cnt)
		return
	}

	fmt.Printf("%s - %d\n", f, cnt)
}

func main() {
	// Задача 9.1
	fruitMarket("киви")
	fruitMarket("яблоки")
	fruitMarket("груши")

	// Задача 9.2
	s := []int{1, 2, 3}
	for _, v := range s {
		cicle := 1
		fmt.Printf("v%d:%d\n", cicle, v)
	STOP_LABEL:
		for _, v := range s {
			cicle := 2
			fmt.Printf("\tv%d:%d\n", cicle, v)
			for _, v := range s {
				cicle := 3
				fmt.Printf("\t\tv%d:%d\n", cicle, v)
				for i, v := range s {
					cicle := 4
					fmt.Printf("\t\t\tv%d:%d\n", cicle, v)
					if i == 1 {
						break STOP_LABEL
					}
				}
			}
		}
	}

	// Задача 9.3
	checkFood("стол")
	checkFood("груша")
	checkFood("яблоко")
	checkFood("апельсин")
	checkFood("тыква")
	checkFood("огурец")
	checkFood("помидор")

}

// Задача 9.3
type foodKind int8
type foodName string

const (
	FRUITS     foodKind = 1
	VEGETABLES foodKind = 2
)

func checkFood(name foodName) {
	foods := map[foodName]foodKind{"апельсин": FRUITS, "яблоко": FRUITS, "сливы": FRUITS, "груша": FRUITS,
		"тыква": VEGETABLES, "огурец": VEGETABLES, "помидор": VEGETABLES}

	kind, ok := foods[name]
	if !ok {
		fmt.Println(name, "- это что-то странное")
		return
	}

	if kind == FRUITS {
		fmt.Printf("%s - это фрукт\n", name)
		return
	}

	fmt.Printf("%s - это овощ\n", name)
}
