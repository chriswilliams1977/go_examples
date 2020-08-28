package arrays

import(
	"fmt"
)

//arrays contain value of same type
//think box of pills with containers for each day
//easy to transport pills in one container
//array = multiple values (elements)in single variable
//cannot grow or shrink = fixed number
//declare by stating number of elements and type
//var myArray [4]string
//zero indexed based
//if values not assigned it contains zero values (0, "", false)
//assign values using array literal
//var myArray [4]string = [4]string{"do","re","me","fa"}
//myArray := [4]string{"do","re","me","fa"}
//myArray := [4]string{
//				"do",
//				"re",
//				"me",
//				"fa",}

//iterating over array using index can be error prone
//if you go higher than index you get error
//can resolve by getting length but this can still cause error if you do <= length as its 0 based
//best to use range as you remove the need for index
func GetArrayElements(){
	notes := [4]string{
		"do",
		"re",
		"me",
		"fa",
	}

	//get the length of notes
	numberOfElements := len(notes)

	//better to use this rather than number in case its not valid
	for i := 0; i < numberOfElements; i++ {
		fmt.Println(notes[i])
	}

	//range is safer and cleaner as no index settings
	//in range first value is index and second element value
	for index, note := range notes {
		fmt.Println(index, note)
	}

	//use blank identifier to get just the index or the value otherwise key value pair returned
	for _, note := range notes {
		fmt.Println(note)
	}
}

