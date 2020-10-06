package interfaces

import (
	"fmt"
	sampleinterfaces "github.com/chriswilliams1977/headfirst/playground/packages/samples/interfaces"
)

//an empty interface accepts any type as it has not method that are required to satisfy it
//if you declare an func that accepts a param with an empty interface as its type
//you can pass it any type

func AcceptAnything(thing interface{}){
	fmt.Println(thing)

	//problem with this is there are no methods to call
	//to call methods on the type you have to use type assertion again
	//easier to write func that accepts concrete type
	tapePlayer, ok := thing.(sampleinterfaces.TapePlayer)
	if ok {
		tapePlayer.Play("Pounding")
	}
}

