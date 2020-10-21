package failures

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetCurrentDirName() string{
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return dir
}

//deferred recover func
func ReportPanic(){
	//set p to interface{} - recover type is an empty interface
	p := recover()
	//if cover returns nil there is no panic
	if p == nil {
		//return out of program
		return
	}
	//panic has occurred
	//check if panic type is error - use type insertion to cast p back to error type
	//else handle random panic if type is not an error
	err, ok := p.(error)
	//if it is an error
	if ok {
		//print error
		fmt.Println(err)
	} else {
		panic(p)
	}
}

//uses recursive function  to scan dir
func ScanDir(path string){
	fmt.Println(path)
	//use ioutil.ReadDir to read directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		//return panic
		//Not ideal as this crashes the program
		//should be reserved for impossible situations
		//use deffer func and recover to allow program to continue executing outside of this func if panic occurs
		panic(err)
	}

	//for each file in directory
	for _, file := range files {
		//join path and the file name
		filePath := filepath.Join(path, file.Name())
		//if file is a directory
		if file.IsDir(){
			//function calls itself and runs through files for that directory
			ScanDir(filePath)
		} else {
			//if its a file print path
			fmt.Println(filePath)
		}
	}
}