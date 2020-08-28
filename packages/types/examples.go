//remember all files in package have same package name
package types

import (
	"fmt"
)

type Animal struct{
	Name string
	Species string
}

func OutputTest(){
	number1 := 5
	fmt.Println("Struct and Types")
	monkey := Animal{
		Name: "Monkey",
		Species: "Mammal",
	}
	fmt.Println(CheckType(number1))
	fmt.Println(CheckType(1.15))
	fmt.Println(CheckType(monkey))

	fmt.Println("Variables")
}