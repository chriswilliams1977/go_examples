package sampleslice

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetArgs(){
	//get everything apart from the first element in the slice
	fmt.Println(os.Args[1:])
}

//Variadic function lets you pass any number of params you want
//use ... before type of last param (has to be last param)
//uses a slice to process this
func VariadicFunctionInRange(min float64, max float64, numbers ...float64){
	var result []float64
	for _,number := range numbers{
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	fmt.Println(result)
}

func GetAveragePurchaseArgs(){
	arguments := os.Args[1:]
	var numbers []float64

	for _, argument := range arguments{
		number, err := strconv.ParseFloat(argument,64)
		if err != nil{
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	//To call varadic func with slice param use ... after your slice
	GetAveragePurchaseVaradic(numbers...)
}

func GetAveragePurchaseVaradic(numbers ...float64){
	var sum float64 = 0

	for _, number := range numbers{
		sum += number
	}

	//need to covert to float64 and len returns int which we cannot use with sum which is a float64
	sampleCount := float64(len(numbers))

	fmt.Printf("Average order amount is %0.2f\n",sum/sampleCount)
}