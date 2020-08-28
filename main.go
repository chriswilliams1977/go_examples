package main

import (
	"github.com/chriswilliams1977/headfirst/playground/packages/samples/arrays"
)

func main() {
	//vars.OutputTest()
	//types.OutputTest()
	//loops.GuessingGame()
	//functions.FormattedValues()

	/*
	totalPaintNeeded, err := functions.PaintCalculator(4.2, 3.0, 10)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Total paint needed is %.2f liters\n", totalPaintNeeded)
	}
	//totalPaintNeeded += functions.PaintCalculator(4.2, 3.0, 10)

	amount := 3
	pointers.PointerSample(&amount)
	pointers.GetPointerTypes()
	pointers.GetPointerAddress(amount)
	pointers.GetPointerValue(amount)

	var myStringPointer = pointers.ReturnPointer()
	fmt.Println(*myStringPointer)

	var myString string = "Test"
	pointers.PrintPointer(&myString)

	greeting.Hi()
	english.SayHello()
	dutch.SayHello()
	arrays.GetArrayElements()
	*/
	samplearray.GetAveragePurchase()
}
