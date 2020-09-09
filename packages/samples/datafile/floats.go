package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//Reads a float from each line of file
func GetFloatsArray(fileName string) ([3]float64,error){
	//declare array to return
	var numbers [3]float64

	//open provided file
	file, err := os.Open(fileName)
	if err != nil{
		//if error opening file return array and error
		return numbers, err
	}

	//var to track array index to assign float from file to
	i := 0

	//returns a bufio.Scanner that reads from file
	scanner := bufio.NewScanner(file)
	//reads each line of file
	for scanner.Scan(){
		//convert file string to float
		numbers[i], err = strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			//if error converting file values return array and error
			return numbers, err
		}
		//move to next index
		i++
	}

	//close file to avoid consuming resources
	err = file.Close()
	if err != nil{
		//if error closing file return array and error
		return numbers, err
	}

	//checks to see if there was an issue with scanner while reading file
	if scanner.Err() != nil{
		////if error scanning file values return array and error
		return numbers, scanner.Err()
	}

	//if we got here we have array populated and ready to return
	return numbers, nil
}

//Reads a float from each line of file
//This now takes a slice not fixed size array so it can grow with datafile entries
func GetFloatsSlice(fileName string) ([]float64,error){
	//declare array to return
	var numbers []float64

	//open provided file
	file, err := os.Open(fileName)
	if err != nil{
		//if error opening file return array and error
		return nil, err
	}

	//returns a bufio.Scanner that reads from file
	scanner := bufio.NewScanner(file)
	//reads each line of file
	for scanner.Scan(){
		//convert file string to float
		number, err := strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			//if error converting file values return array and error
			return nil, err
		}
		//append new number to slice numbers
		numbers = append(numbers, number)
	}

	//close file to avoid consuming resources
	err = file.Close()
	if err != nil{
		//if error closing file return array and error
		return nil, err
	}

	//checks to see if there was an issue with scanner while reading file
	if scanner.Err() != nil{
		////if error scanning file values return array and error
		return numbers, scanner.Err()
	}

	//if we got here we have array populated and ready to return
	return numbers, nil
}