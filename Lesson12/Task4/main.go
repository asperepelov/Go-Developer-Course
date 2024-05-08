package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	//4-й вариант. Проверка указателя и определение поведения по умолчанию
	if d == nil {
		return DefaultSing()
	}
	return d.voice
}

// 4-й вариант. Определение поведения по умолчанию
func DefaultSing() string {
	return "кря для не созданного экземпляра"
}

func main() {
	//var d *Duck
	var d *Duck = new(Duck) // 1-й вариант. Указателю присвоить созданный экземпляр Duck
	d.voice = "Кря"
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)

	//2-й вариант. создавать переменную типа Duck и передавать указатель на неё в Sing
	var d1 = Duck{voice: "кря-кря"}
	song, err = Sing(&d1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)

	//3-й вариант. во избежание путаницы указывать получателя без указателя
	//это вынудит явно создавать экземпляр
	//func (d Duck) Sing() string

	// 4-й вариант. Определение поведения без использования поля экземпляра
	var d4 *Duck
	song, err = Sing(d4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
