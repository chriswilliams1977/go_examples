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

//uses recursive function  to scan dir
func ScanDir(path string) error{
	fmt.Println(path)
	//use ioutil.ReadDir to read directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	//for each file in directory
	for _, file := range files {
		//join path and the file name
		filePath := filepath.Join(path, file.Name())
		//if file is a directory
		if file.IsDir(){
			//function calls itself and runs through files for that directory
			err := ScanDir(filePath)
			if err != nil{
				return err
			}
		} else {
			//if its a file print path
			fmt.Println(filePath)
		}
	}
	return nil
}