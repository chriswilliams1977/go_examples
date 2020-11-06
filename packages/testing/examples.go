package testing

import "strings"

//example func to take slice of strings and join them
//same as fmt.Println(strings.Join([]string{"state","of","the","art"},"-"))
//produces state-of-the-art
//if we run the below it works ok for two words but how do we test this
func JoinWithCommas(phrases []string) string{
	if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) ==2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		//strings.join takes a slice of strings and a string to join
		//:len(phrases)-1] gets all strings upto last string and joins them with comma in between
		result := strings.Join(phrases[:len(phrases)-1],", ")
		//add and
		result += " and "
		//add the last string in slice
		result += phrases[len(phrases)-1]
		//returns string with items joined
		return result
	}
}