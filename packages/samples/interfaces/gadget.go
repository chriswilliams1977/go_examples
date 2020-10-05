package sampleinterfaces

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

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

func (tr TapeRecorder) Record(){
	fmt.Println("Recording")
}

//play songs on a device
//if you want to use either a TapePlayer or TapeRecorder to play songs you need an interface
//all you need is a type that can both play and stop
func PlayList(device Player, songs []string){
	for _,song := range songs{
		//needs a type that can play
		device.Play(song)
	}
	//needs a type that can stop
	device.Stop()
}

//This method uses type assertion to call method of concrete type
func TryOut(player Player){
	player.Play("Test Track")
	player.Stop()
	//type assertion have a second optional return value
	//indicates whether assertion was successful
	var p Player = TapeRecorder{}
	recorder, ok := p.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not TapeRecorder")
	}
}