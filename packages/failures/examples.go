package failures

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//open a file
//takes filename and returns pointer to a file and an error
func OpenFile(filename string) (*os.File, error){
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

//close a file
//takes pointer to a file
func CloseFile(file *os.File){
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(filename string)([]float64, error){
	//create slice for numbers
	var numbers []float64
	//open file
	file, err := OpenFile(filename)
	//handle error
	if err != nil {
		return nil, err
	}

	//close file
	//this ensures it will be called even if there is an error parsing the contents
	defer CloseFile(file)

	//scan file
	scanner := bufio.NewScanner(file)
	//for each row scanned
	for scanner.Scan(){
		//convert file text to float
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		//handle error
		if err != nil {
			return nil, err
		}
		//add float to numbers slice
		numbers = append(numbers,number)
	}

	//handle scanner error if it occurred as it fails silently
	if scanner.Err() != nil{
		return nil, scanner.Err()
	}
	//return slice of floats
	return numbers, nil
}

func DeferExample() error{
	//defers call until all other code is executed
	defer fmt.Println("Goodbye!!")
	fmt.Println("Hello")
	//return error
	return fmt.Errorf("I dont want to talk!!")
	fmt.Println("Nice weather hey!!")
	return nil
}