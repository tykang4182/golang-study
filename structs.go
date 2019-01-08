package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Daniel", 29})
	fmt.Println(person{name: "Cherie", age: 31})
	fmt.Println(person{name: "Alesso"})
	fmt.Println(&person{name: "Pat", age: 40})

	s := person{name: "Barry", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
