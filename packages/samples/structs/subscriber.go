package samplestructs

type Subscriber struct{
	Name string
	Rate float64
	Active bool
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
	s.Name = "changed name"
}
