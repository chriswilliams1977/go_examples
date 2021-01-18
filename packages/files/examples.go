package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func ReadFile(){
	file, err := os.OpenFile("./packages/files/aardvark.txt",os.O_RDONLY,os.FileMode(0600))

	check(err)

	//this is ok when reading but not when writing
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	check(scanner.Err())
}

func WriteFile(){
	//bitwise operators are used to compare binary numbers
	//bitwise & OR checks if the bits either side are 1 or 0
	//openfile uses bitwise OR to append constants for the second param
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	//filemode determines permissions users has with file
	//three sets of permissions: owner of file, users file is assigned too, users on system
	//permission are rwx - read, write, execute
	//Filemode uses octal notation - eight digits 0 - 7
	//Octal notation is the 3 digits proceeding a 0 i.e 0100 or 0111 or 0001
	//at position 7 it resets e.g 0000 to 0007 then 0010 to 0017
	//octal number can be represented by 3 bits of memory
	//this is the exact amount needed to store permissions
	//used by Unix chmod
	//decimal notation is 0 - 10
	//binary is 0 and 1
	file, err := os.OpenFile("./packages/files/aardvark.txt",options,os.FileMode(0600))

	check(err)

	_, err = file.Write([]byte("amazing!\n"))

	check(err)

	err = file.Close()

	check(err)
}

func BitwiseExample(){
	//bitwise operators work with bit representation of a value
	a := 10 //1010
	b := 3  //0011

	//checks if both bits are 1  = 0010 = 2
	fmt.Printf("%v\n", a&b)
	//checks if a bit is 1 = 1011 = 2
	fmt.Printf("%v\n", a|b)
	//checks if bit is 1 but not both = 1001 = 9
	fmt.Printf("%v\n", a^b)
	//checks if no bits are 1  = 0100 = 8
	fmt.Printf("%v\n", a&^b)

}

//A bit shift moves each digit in a number's binary representation left or right.
//There are three main types of shifts:
//Left Shift: When shifting left, the most-significant bit is lost, and a 00 bit is inserted on the other end.
//0010 << 1  →  0100
//0010 << 2  →  1000
//A single left shift multiplies a binary number by 2:   0010 << 1  →  0100
//0010 is 2
//0100 is 4
//Logical Right: When shifting right with a logical right shift, the least-significant bit is lost and a 00 is inserted on the other end.
//1011 >>> 1  →  0101
//1011 >>> 3  →  0001
func BitShiftingExample(){

}