package maps

import (
	"fmt"
	"sort"
)

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

//in this example you can ask for the grade of carl even though he is not logged due to zero values
//map key optionally returns second boolean value which states if value has been assigned
//or false if zero value
//devs use ok  to check if assigned
func Status(name string){
	//map literal with grades
	grades := map[string]float64{"Alma":0,"Rohit":86.5}
	//get map value of student with name
	grade, ok := grades[name]
	if !ok{
		fmt.Println("no grade recorded")
	} else {
		//if value less than 60
		if grade < 60 {
			fmt.Printf("%s is failing\n",name)
		} else {
			fmt.Printf("%s passed\n",name)
		}
	}

}

func CheckZeroValue(){
	//create a slice with letters
	data := []string{"a","c","e","a","e"}
	//create a map to count numbers of each letter
	counts := make(map[string]int)
	//loop over letters
	for _,item := range data{
		//for each letter found increment map value (int)
		counts[item]++
	}
	//make second slice of letters
	letters := []string{"a","b","c","d","e"}
	//loops over letters
	for _, letter := range letters{
		//for each letter check if it is assigned value in counts maps
		count, ok := counts[letter]
		//if not assigned value
		if !ok {
			fmt.Printf("%s: not found\n",letter)
		}else{
			fmt.Printf("%s: %d\n",letter,count)
		}
	}
}

func DeleteMapKey(){
	//create var for zero check
	var ok bool
	//create map to hold ranks
	ranks := make(map[string]int)
	//create var to get rank
	var rank int
	//set rank bronze to value 3
	ranks["bronze"] = 3
	//get rank value for bronze and set rank to true/false if value assigned
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n",rank, ok)

	//delete rank bronze
	delete(ranks, "bronze")
	//get rank value for bronze and set rank to true/false if value assigned
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n",rank, ok)
}

func UnorderedMapLoop(){
	grades := map[string]float64{"Alma":74.3,"Rohit":86.5,"Carl":59.7}
	//for loop processes keys and values in random order because map is unordered
	//emit value  if you do not want to print key
	//use _, grade if you only want value
	for name, grade := range grades{
		fmt.Printf("%s has a grade of %0.1f%%\n",name,grade)
	}
}

func OrderedMapLoop(){
	grades := map[string]float64{"Alma":74.3,"Rohit":86.5,"Carl":59.7}
	//to order map keys and values you need to create a slice and add the names
	//then sort alphabetically
	//print names from slice
	var names []string
	for name := range grades{
		names = append(names,name)
	}
	//alphabetical sorting
	sort.Strings(names)
	//print names in slice
	for _, name := range names{
		fmt.Printf("%s has a grade of %0.1f%%\n",name,grades[name])
	}
}