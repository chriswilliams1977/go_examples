package sampleslice


import (
	"fmt"
	"github.com/chriswilliams1977/headfirst/playground/packages/samples/datafile"
	"log"
)

func GetAveragePurchase(){
	numbers, err := datafile.GetFloatsSlice("data.txt")
	if err != nil{
		log.Fatal(err)
	}

	sum := 0.0

	for _, order := range numbers{
		sum += order
	}

	//need to covert to float64 and len returns int which we cannot use with sum which is a float64
	sampleCount := float64(len(numbers))

	fmt.Printf("Average order amount is %0.2f\n",sum/sampleCount)
}