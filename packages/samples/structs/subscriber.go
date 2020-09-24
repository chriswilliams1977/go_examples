package samplestructs

import "fmt"

//must be capitals to be exported
//same goes for the fields
type Subscriber struct{
	Name string
	Rate float64
	Active bool
	//annonymous field = remove name and just have type
	//can use field type name as if it was the name of field
	//referred to as being embedded in outer struct
	//these are promoted and can be referenced as they belong to outer struct
	//thus you can do subscriber.City
	Address
}

type Employee struct{
	Name string
	Salary float64
	Address
}

type Address struct{
	Street string
	City string
	State string
	PostalCode string
}

//takes a pointer to a subcriber and prints properties of subscriber pointer
func PrintInfo( s *Subscriber){
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?:", s.Active)
}

//returns a pointer to a subscriber
func DefaultSubscriber(name string) *Subscriber{
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

//applies discount to pointer to subscriber
func ApplyDiscount(s *Subscriber){
	s.Rate = 4.99
}

//remember go is pass-by-value
//if you create a subscriber locally and then pass to a function that is not using pointers
//the func will modify local version not the one passed thus use pointers
func NewWithOutPointer(s Subscriber) {
	s.Name = "changed name"
}

//use * to update value of pointer
//* = value of struct
func NewWithPointer(s *Subscriber){
	//with structs you dont need to add * in front of var
	s.Name = "changed name"
}
