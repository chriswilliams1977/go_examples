//sample file how to read files
package samplearray

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

func ReadFile(){
	//caller function allows you to get
	//params are pc, file, line, ok
	//you provide int which tells which executing func to skip
	_, filedir, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filedir), "data.txt")

	//return pointer to an os.File and an error
	file, err := os.Open(filepath)
	if err != nil{
		log.Fatal(err)
	}

	//returns a bufio.Scanner that reads from file
	scanner := bufio.NewScanner(file)
	//reads each line of file
	for scanner.Scan(){
		//returns string representing data read
		fmt.Println(scanner.Text())
	}

	//close file to avoid consuming resources
	err = file.Close()
	if err != nil{
		log.Fatal(err)
	}

	//checks to see if there was an issue with scanner while reading file
	if scanner.Err() != nil{
		log.Fatal(scanner.Err())
	}
}