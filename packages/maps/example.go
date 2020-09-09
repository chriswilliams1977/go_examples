package maps

import "fmt"

//map with unknown values to be added
func CreateMap(){
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
}

//map with known values to be added
func CreateMapLiterals(){
	elements := map[string]string{
		"H":"Hydrogen",
		"Li":"Lithium",
	}
	fmt.Println(elements["Li"])

	//if empty
	//emptyMap = map[string]string{}
}