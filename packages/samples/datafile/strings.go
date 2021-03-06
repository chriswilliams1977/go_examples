//Reads data from a file
package datafile

import (
	"bufio"
	"os"
)

func GetStrings(fileName string) ([]string,error){
	//declare slice to hold values
	var lines []string

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
		//Get line of text
		line := scanner.Text()
		//append text to slice
		lines = append(lines, line)
	}

	//close file to avoid consuming resources
	err = file.Close()
	if err != nil{
		//if error closing file return nil and error
		return nil, err
	}

	//checks to see if there was an issue with scanner while reading file
	if scanner.Err() != nil{
		////if error scanning file values return nil and error
		return nil, scanner.Err()
	}

	//if we got here we have slice populated and ready to return
	return lines, nil
}