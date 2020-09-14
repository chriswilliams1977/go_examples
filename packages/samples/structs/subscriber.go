package samplestructs

import (
	"fmt"
)

var subscriber struct{
	name string
	rate float64
	active bool
}

func SetSubscriber(){
	subscriber.name = "bob"
	subscriber.rate = 4.99
	subscriber.active = true

	fmt.Println("Name: ", subscriber.name)
	fmt.Println("Rate: ", subscriber.rate)
	fmt.Println("Active: ", subscriber.active)
}