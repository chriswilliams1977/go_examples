package main

import (
	d "github.com/chriswilliams1977/headfirst/playground/packages/definedtype"
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

	samplearray.GetAveragePurchase()
	sampleslice.GetAveragePurchase()
	sampleslice.GetArgs()
	sampleslice.VariadicFunctionInRange(1,100,-12.5, 3.2, 0, 50, 103.5)
	sampleslice.GetAveragePurchaseArgs()

	samplemaps.GetDataUsingSlice()
	maps.CreateMap()
	maps.Status("Alma")
	maps.Status("Rohit")
	maps.Status("Carl")
	maps.CheckZeroValue()
	maps.DeleteMapKey()
	samplemaps.GetDataUsingMap()
	*/
	//structs.GetStruct()
	//Pass struct and amend name locally without using a pointer
	//subnopointer := s.Subscriber{"Chris",2.20,false}
	//s.NewWithOutPointer(subnopointer)
	//fmt.Println(subnopointer)
	//Pass struct and amend name locally without using a pointer
	//subwithpointer := s.Subscriber{"Bob",3.0,true}
	//pass address of pointer using &
	//&  = address of struct
	//s.NewWithPointer(&subwithpointer)
	//fmt.Println(subwithpointer)
	/*
	subscriber1 := s.DefaultSubscriber("Aman Singh")
	s.ApplyDiscount(subscriber1)
	s.PrintInfo(subscriber1)

	subscriber2 := s.DefaultSubscriber("Beth Ryan")
	s.PrintInfo(subscriber2)
	structs.GetStructLiteral()

	subscriber := s.Subscriber{Name:"Bob"}
	//address is an anonymous type
	//address is promoted to outer struct
	//so you can reference address fields from outer struct
	subscriber.Street = "Rydons Lane"
	subscriber.City = "Croydon"
	subscriber.State = "Surrey"
	subscriber.PostalCode = "CR5 1SU"
	fmt.Println(subscriber)
	*/
	d.GetFuel()
	g := d.Gallons(1.2)
	g.GallonsReceiverWithParams(10, true)
	g.GallonsReceiverMultiReturns()
	d.GetMyTypeValues()
}
