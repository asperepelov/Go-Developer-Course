package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	msg    string
	status int8
}

func (e myFirstError) Error() string {
	return fmt.Sprintf("%d: %s", e.status, e.msg)
}

func main() {
	// Задача 11.1
	err1 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", errors.New("ошибка1")))
	fmt.Println(err1)

	// Задача 11.2
	err2 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", errors.New("ошибка1")))
	err21 := errors.Unwrap(err2)
	fmt.Println(err21)

	// Задача 11.3
	err31 := errors.New("ошибка1")
	err3 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", err31))
	fmt.Println(errors.Is(err3, err31))

	// Задача 11.4
	err4 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", errors.New("ошибка1")))
	if !errors.As(err4, new(myFirstError)) {
		fmt.Println("В созданной цепочке нет ошибки типа myFirstError")
	} else {
		fmt.Println("В созданной цепочке есть ошибка типа myFirstError")
	}

}
