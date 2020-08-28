package loops

import (
	"bufio"
	"fmt"
	"github.com/chriswilliams1977/headfirst/playground/packages/keyboard"
	"math/rand"
	"os"
	"log"
	"strconv"
	"strings"
	"time"

)

func OutputTest() {
	//Input func moved top reusable package
	//Creates use prompt. Println does \n after but Print does not skip to new cmd line
	//bufio reader receives user input from Stdin
	//All keyboard inputs go to Stdin
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil{
		log.Fatal(err)
	}


	var status string
	if grade == 100 {
		status = "a top score, well done!!"
		fmt.Println()
	} else if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is",status)
}

//condition statement
//pass if >60%, fail if < 60%
func Result(){

}

func ErrorExample() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}

func GuessingGame(){
	//To get different random numbers we need to use seed
	//Seed mean give it a value to generate other random numbers
	//But is you keep giving same seed you will get same random number
	//Unix time is number of seconds since jan 1 1970
	//generate random number between 1-100 hence + 1
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	success := false
	for guesses := 0 ; guesses < 10; guesses++ {

		if guesses == 0 {
			//ask player what number they think it is
			fmt.Print("What number am I thinking of?")
		} else {
			fmt.Println("You have", 10-guesses, "guesses left")
		}

		//Get players entry
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		var feedback string
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		numberGuess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}
		if numberGuess < target {
			feedback = "Opps, your guess was too LOW"
		} else if numberGuess > target {
			feedback = "Opps, your guess was too HIGH"
		}else {
			success = true
			feedback = "Good job, you guessed it!!"
			fmt.Println(feedback)
			break
		}
		fmt.Println(feedback)
	}
	//make an if because is player wins and breaks out of loop dont want to show this message!
	if !success {
		fmt.Println("Bad luck you ran out of guesses. The number was", target)
	}

}

func OperatorsTest(){
	for x := 1; x <= 5; x += 2{
		fmt.Println(x)
	}
}