package webapp

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string){
	//writer.Write does not accept strings so we need to convert string to byte
	//byte type is used to hold raw data
	//need to convert to string to get readable text
	//add text to the response
	_,err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

//Handler
//writer holds response to send to the browser
//request represents the request from the browser
func EnglishHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Hello, web!")
}

func FrenchHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Salut, web!")
}

func HindiHandler(writer http.ResponseWriter, request *http.Request){
	write(writer,"Nasmaste, web!")
}