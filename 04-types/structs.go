package types

import (
	"fmt"

	
)

// type untuk pembuatan struct
type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name:name}
	p.age = 42
	return &p
}

func GoStruct(){
	// langsung memanggil struct 
	fmt.Println(person{"bob", 20})
	
	fmt.Println(person{name: "Alice", age: 30})
	//memanggil dengan alamat
	fmt.Println(&person{name: "ann", age :30})
	// memanggil func newPerson
	fmt.Println(newPerson("jon"))

	s:= person{name:"sean", age:50}
	fmt.Println(s)
	sp:= &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	// anynomous struct, hanya dengan ditampung dalam variable hanya untuk sekali pakai
	dog:= struct{
		name string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}