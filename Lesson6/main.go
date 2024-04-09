package main

import "fmt"

type Contract struct {
	ID     int
	Number string
	Date   string
}

func (c Contract) ToStr() string {
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}

func main() {
	// Задача 6.1
	c1 := Contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v\n", c1)

	// Задача 6.2
	type contract struct {
		ID     int
		Number string
		Date   string
	}
	c2 := contract{
		ID: 1, Number: "#000A101\t01", Date: "2024-01-31",
	}
	fmt.Printf("%+v\n", c2)

	// Задача 6.3
	c3 := Contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}
	fmt.Println(c3.ToStr())

	// Задача 6.4
	type contactInfo struct {
		Addss string
		Phone string
	}
	type user struct {
		ID   int
		Name string
		contactInfo
	}
	type employee struct {
		ID   int
		Name string
		contactInfo
	}
	u := user{
		ID:   1,
		Name: "Mark",
		contactInfo: contactInfo{
			Addss: "NN",
			Phone: "111112234",
		},
	}
	e := employee{
		ID:   1,
		Name: "Tom",
		contactInfo: contactInfo{
			Addss: "Mosc",
			Phone: "44412234",
		},
	}
	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)

}
