package samplemaps

import (
	"log"
	"fmt"
	"github.com/chriswilliams1977/headfirst/playground/packages/samples/datafile"
)

func GetDataUsingSlice(){
	//get lines from datafile
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil{
		log.Fatal(err)
	}
	//create two slices, one to hold names and one to hold counts
	var names []string
	var counts []int
	//iterate over each line in file
	for _, line := range lines{
		matched := false
		//iterate over names slices
		for i, name := range names {
			//if name in file matches name in slice
			if name == line{
				//increment count slice
				counts[i]++
				matched = true
			}
		}
		//if no match
		if matched == false{
			//add name to name slice
			names = append(names, line)
			//set slice count to 1
			counts = append(counts, 1)
		}
	}
	//print names and count
	for i, name := range names{
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

func GetDataUsingMap(){
	//get lines from datafile
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil{
		log.Fatal(err)
	}
	//create map
	//iterate over map and increment vote
	//do not need zero value check as lines exist in file
	counts := make(map[string]int)
	//iterate over each line in file
	for _, line := range lines{
		counts[line]++
	}

	for name, count := range counts{
		fmt.Printf("Votes for %s: %d\n",name, count)
	}
}