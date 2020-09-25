package definedtype

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64
type MyType string

//example of receiver method with params
func (g Gallons) GallonsReceiverWithParams(number int, flag bool){
	fmt.Println(g)
	fmt.Println(number)
	fmt.Println(flag)
}

//example of receiver method with multiple returns
func (g Gallons) GallonsReceiverMultiReturns() int{
	return int(g)
}

//Example to show pointer reference on defined types
//this method takes a value receiver for the defined type
func (m MyType) method(){
	fmt.Println("Method with value receiver")
}

//this method takes a pointer receiver for the defined type
func (m *MyType) pointerMethod(){
	fmt.Println("Method with pointer receiver")
}

func GetMyTypeValues(){
	//note you must store values in a var for go to automatically convert
	//you cannot do &MyType("a value") or MyType("a value").pointerMethod
	//same goes for when referencing a pointer as a var
	//create var of type MyType
	value := MyType("a value")
	//create a var that is a pointer to value
	pointer := &value
	//call method with local value
	value.method()
	//notice we can call pointer method with local value and go converts to pointer automatically - dont need to convert
	value.pointerMethod()
	//go automatically retrieves value at pointer and not local value - dont need to convert
	pointer.method()
	pointer.pointerMethod()
}

func (g Gallons) ToLiters() Liters{
	//can convert as same underlying type
	return Liters(g * 3.785)
}

//Note you can have same name methods as longs as they are for different types
func (l Liters) ToGallons() Gallons{
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons{
	return Gallons(m * 0.000264)
}

func (g Gallons) ToMilliliters() Milliliters{
	return Milliliters(g * 3785.41)
}


func GetFuel(){
	//runs on gallons with 1.2 left in the tank
	carfuel := Gallons(1.2)
	//runs on liters with 4.5 left in the tank
	busfuel := Liters(4.5)

	soda := Milliliters(500)

	//see how many gallons I have if I add 40 liters
	carfuel += busfuel.ToGallons()
	//see how many liters I have if I add 30 gallons
	busfuel += carfuel.ToLiters()
	fmt.Println("Car fuel: %0.1f gallons", carfuel)
	fmt.Println("Bus fuel: %0.1f liters", busfuel)

	fmt.Printf("%0.3f gallons equals %0.3f liters\n", carfuel, carfuel.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", carfuel, carfuel.ToMilliliters())
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", soda, soda.ToGallons())
}