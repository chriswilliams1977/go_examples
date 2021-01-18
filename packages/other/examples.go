package other

import (
	"io/ioutil"
	"log"
	"time"
)

//like for loops you can intialize a var for use in the statement
//scope is limited to the if statement
func ExampleIfErrCondition(){
	/*
	if count := 5; count > 4 {
		fmt.Println("count is", count)
	}*/

	//err declared twice but has different scopes
	if err := saveString("english.txt","Hello"); err != nil{
		log.Fatal(err)
	}

	if err := saveString("hindi.txt","Namaste"); err != nil{
		log.Fatal(err)
	}
}

func saveString(filename string, str string) error{
	err := ioutil.WriteFile(filename, []byte(str), 0600)
	return err
}

//runes are used to store UTF8 chars
//rune is 1-4 bytes - can store chars that need > 1 byte
//ASCII chares are only 1 byte for example english chars

//buffered channels blocks channel until another goroutine recieves value
//buffered channels lets you add values to a buffer in the channel
//these fill up until limit reached
//for example channel := make(chan string, 3) - 3 values can be added to buffer
//when goroutine recieves values from channel it takes earliest added value from buffer
func SendLetters(channel chan string){
	//send 4 letters sleeping 1 second before each value
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"
}