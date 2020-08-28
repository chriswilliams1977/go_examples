package pointers

import (
	"fmt"
	"reflect"
)

func PointerSample(number *int) {
	*number *= 2
	fmt.Println(number)
}

func GetPointerTypes(){
	//pointer can only hold one type in this case an int
	var myInt int
	myIntPointer := &myInt
	fmt.Println(reflect.TypeOf(&myIntPointer))

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
}

func GetPointerAddress(number int){
	//pointer can only hold one type in this case an int
	myInt := number
	fmt.Println(&myInt)
}

func GetPointerValue(number int){
	//pointer can only hold one type in this case an int
	myInt := number
	myIntPointer := &myInt
	fmt.Println(*myIntPointer)

	*myIntPointer = 8
	fmt.Println(*myIntPointer)
}

func ReturnPointer() *string{
	myString := "I am a pointer"
	return &myString
}

func PrintPointer(myPointer *string){
	fmt.Println(*myPointer)
}