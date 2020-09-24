package definedtype

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

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

func ToLiters(g Gallons) Liters{
	//can convert as same underlying type
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons{
	return Gallons(l * 0.264)
}

func GetFuel(){
	//runs on gallons with 1.2 left in the tank
	carfuel := Gallons(1.2)
	//runs on liters with 4.5 left in the tank
	busfuel := Liters(4.5)

	//see how many gallons I have if I add 40 liters
	carfuel += ToGallons(Liters(40.0))
	//see how many liters I have if I add 30 gallons
	busfuel += ToLiters(Gallons(30.0))
	fmt.Println("Car fuel: %0.1f gallons", carfuel)
	fmt.Println("Bus fuel: %0.1f liters", busfuel)
}