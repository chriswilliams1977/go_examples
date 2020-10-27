package goroutines

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPageWeight(url string){
	fmt.Println("Getting",url)

	//http.Get() returns http response obj
	//this obj is a struct with a body that is content of the page
	//body satisfies the the io package readclose interface - can read and close connection
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	//close the connection after we have read the page
	defer resp.Body.Close()

	//read contents of page
	//returns a slice of bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//convert bytes into string
	fmt.Println(len(body))
}

func ChannelExample(myChannel chan string){
	myChannel <- "hi"
}

//two functions which recieve channels as part of goroutines
func ChannelABC(channel chan string){
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func ChannelDEF(channel chan string){
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

//example of blocking slowed down
func ReportNap(name string, delay int){
	//fo every second goroutine is asleep print message sleeping
	for i := 0; i < delay; i++{
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	//wake up goroutine message
	fmt.Println(name, "wakes up!")
}

//func that cause goroutine to nap
//recieve a channel
func SendGoRoutine(myChannel chan string){
	//sleep goroutine for 2 seconds
	//every second goroutine is asleep it prints a sleeping message
	ReportNap("sending goroutine",2)

	fmt.Println("***sending value***")
	//send 2 value to channel
	myChannel <- "a"
	fmt.Println("***sending value***")
	//send value to channel
	myChannel <- "b"
}