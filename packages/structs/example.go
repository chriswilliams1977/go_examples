package structs

import "fmt"

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