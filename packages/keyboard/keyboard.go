package keyboard

import(
	"log"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error){
	reader := bufio.NewReader(os.Stdin)
	//this actually gets user input value, with a rune to mark end of input '\n'
	//ReadString return a second values, an error if unable to read string
	//You must handle this by using blank identifier _ or log the error
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//you need this otherwise \n is the value read not the number
	//Alternative is
	////scanner := bufio.NewScanner(os.Stdin)
	//fmt.Print("Enter a number")
	//scanner.Scan()
	//input := scanner.Text()
	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input,64)
	if err != nil {
		log.Fatal(err)
	}
	return number, nil
}