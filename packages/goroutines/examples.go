package goroutines

import (
	"fmt"
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