package webapp

import (
	"fmt"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string){
	//writer.Write does not accept strings so we need to convert string to byte
	//byte type is used to hold raw data
	//need to convert to string to get readable text
	//add text to the response
	_,err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

//Handler
//writer holds response to send to the browser
//request represents the request from the browser
func EnglishHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Hello, web!")
}

func FrenchHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Salut, web!")
}

func HindiHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Nasmaste, web!")
}

//func as first class citizen example
//means func can be assigned to a var
//and then used as that var
func sayHi(){
	fmt.Println("Hi")
}

func sayBye(){
	fmt.Println("Bye")
}

func divide(a int, b int) float64{
	return float64(a)/float64(b)
}

func multiply(a int, b int) float64{
	return float64(a * b)
}

func doMath(passedFunc func(int,int) float64){
	result := passedFunc(10,2)
	fmt.Println(result)
}

func FuncVarExample(){
	//example assigning func to a var
	var myFunction func()
	myFunction = sayHi
	myFunction()

	//example calling a func and passing a func as a param
	FuncParamExample(sayHi)
	FuncParamExample(sayBye)

	//example of passing a fun with specific type (params and return value)
	//define var as func with as a type that expects int params and returns float64
	var mathFunc func(int, int) float64
	//assign divide to my var as it matches the type definition
	mathFunc = divide
	//print response
	fmt.Println(mathFunc(5,2))
	//if we did mathFunc  = sayHi it would error as type definition wrong

	doMath(divide)
	doMath(multiply)
}

func FuncParamExample(theFunction func()){
	theFunction()
	theFunction()
}



