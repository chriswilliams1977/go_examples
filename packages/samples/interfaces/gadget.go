package sampleinterfaces

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (tp TapePlayer) Play(song string){
	fmt.Println("Playing", song)
}

func (tp TapePlayer) Stop(){
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	Microphones int
}

func (tr TapeRecorder) Play(song string){
	fmt.Println("Playing", song)
}

func (tr TapeRecorder) Stop(){
	fmt.Println("Stopped")
}

//play songs on a device
//if you want to use either a TapePlayer or TapeRecorder to play songs you need an interface
//all you need is a type that can both play and stop
func PlayList(device TapePlayer, songs []string){
	for _,song := range songs{
		//needs a type that can play
		device.Play(song)
	}
	//needs a type that can stop
	device.Stop()
}