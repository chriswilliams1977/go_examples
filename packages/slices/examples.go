package slices

import "fmt"

//Slices are built on top of arrays
//Slice is a view into the array
//Like a microscope showing specific part of an array
func DeclareSlice(){
	underlyingArray := [5]string{"a","b","c","d","e"}
	//slice starts at a and ends at d but does not include d
	//you can omit start and stop which will default to 0 and end of the slice
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1[1])

	//Slices can have extra elements added to them, arrays are fixed length
	//adding element changes underlying array - using append
	//if underlying array is too small to append value and new array is created and values copied over
	//so assign return value back to same slice


	//usually easier to use make and literals to manage slice rather than create the array
	//declaring slice does not automatically create one you must use make
	//declares slice
	//var notes []string
	//declares and creates slice, now you can assign values
	//notes := make([]string, 7)
	//Literal is where you initialise slice with values (same goes for array), make you are not assigning values yet
	noteslit := []string{"do","re","mi"}
	fmt.Println(noteslit[2])
}
