package structs

import "fmt"

//type definition
type person struct{
	firstname string
	lastname string
	age int
	height float64
	weight float64
}

func GetStruct(){

	var chris person
	chris.firstname = "Chris"
	chris.lastname = "Williams"
	chris.age = 43
	chris.height = 181
	chris.weight = 83.5
	fmt.Printf("%#v",chris)
}

func GetStructLiteral(){

	pat := person{firstname:"Pat",lastname:"Williams",height:150}
	fmt.Printf("%#v",pat)
}