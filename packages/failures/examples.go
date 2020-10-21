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

func RecursionExample(){
	count(1,3)
}

//example of function recursion
func count(start int, end int){
	fmt.Printf("count(%d, %d) called \n", start, end)
	fmt.Println(start)
	if start < end {
		//function calls itself
		count(start+1, end)
	}
	fmt.Printf("Returning from count (%d, %d) call \n", start, end)
}

func PanicExample(){
	//if func is panicing this deferred recover will always get completed a recover from panic
	defer RecoverExample()
	panic("oh no")
	//this never gets called as panic stops func executing
	//recover()
	fmt.Println("I wont be run")
}

func RecoverExample(){
	recover()
}

func PanicTest(){
	PanicExample()
	//this gets run as recover is called
	//fun PanicExample func returns wtihout calling print
	fmt.Println("Exiting normally")
}