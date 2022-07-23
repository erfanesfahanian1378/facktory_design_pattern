package main

import "log"

func main()  {
	//p := NewPerson(1999)
	//log.Print(p)
	farm(32)
}

type Animal interface {
	Sound() string
}

type Cat struct {
}

func NewCat() Cat  {
	return Cat{}
}

func (c Cat) Sound() string {
	return "Meow"
}

type Dog struct {
}

func NewDog() *Dog {
	return &Dog{}
}

func (d Dog) Sound() string  {
	return "Hap"
}

func farm(x int) {
	var a Animal
	if x > 42 {
		a = NewCat()
	}else {
		a = NewDog()
	}

	log.Print(a.Sound())
}

//type Person struct {
//	age int
//}
//
//func NewPerson(yearsOfBirth int) Person  {
//	return Person{age: 2022 - yearsOfBirth}
//}