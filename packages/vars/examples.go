package vars

import (
	arithmetic "github.com/chriswilliams1977/headfirst/playground/packages/numbers"
	"math"
	"strings"
	"fmt"
)

type Animal struct{
	Name string
	Species string
}

func OutputTest() {
	//When using modules
	//import module path + package folder name
	//then call package + func
	//If you want to import remote package push to git with version
	//then go get git repo
	//check pkg/mod repo for package

	//rune stores data as numeric code not the text characters
	runetest := 'A'
	chartest := "A"
	fmt.Println("Rune",runetest)
	fmt.Println("Character",chartest)
	fmt.Println(strings.Title("to upper case"))
	fmt.Println(math.Floor(1.6))

	fmt.Println("Numbers")
	number1 := 5
	number2 := 5

	fmt.Println(arithmetic.Addition(number1,number2))

	//Use this declaration when you do not know the value of a variable
	var myZeroValueInt int
	var myAssignedInt int = 7
	var myMultipleValuesA, myMultipleValuesB = 1.2, 4.9
	fmt.Println(myZeroValueInt)
	fmt.Println(myAssignedInt)
	fmt.Println(myMultipleValuesA,myMultipleValuesB)

	//if you know the value straight away use this declaration
	//myValue := 7

	playVarsExample()

	fmt.Println("Conversions")
	length := 1.2
	width := 2
	//to convert add the type then the var
	fmt.Println("Area is", length * float64(width))

	playConversionExample()
}

func playVarsExample(){
	var originalCount int  = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")

	eatenCount = originalCount-eatenCount

	fmt.Println("There are", eatenCount, "apples left.")

}

func playConversionExample(){
	var price int = 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax,"dollars")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is",total,"dollars")

	var availableFunds int = 20
	fmt.Println(availableFunds,"dollars available")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}