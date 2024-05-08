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
	return d.voice
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

	//3-й вариант. во избежание путаницы указывать получателя  без указателя
	//func (d Duck) Sing() string

}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
