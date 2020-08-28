package functions

import "fmt"

func PaintCalculator(width, height, surfacePerLitre float64) (float64,error){
	if width < 0 {
		//returns from function if there is a problem with an error
		return 0, fmt.Errorf("a width of %0.2f is invalid",width)
	}
	if height < 0 {
		return 0,fmt.Errorf("a height of %0.2f is invalid",height)
	}
	area := width * height
	//Each litre of paint covers 10 square meters
	//Can use %.2f to round to two decimal places
	fmt.Printf("%.2f liters needed\n", area/surfacePerLitre)
	return area/surfacePerLitre, nil
}

func SayHelloToGroup(line string, times int){
	for i := 0; i < times; i++ {
		fmt.Println("Hello "+ line)
	}
}

//Please read this comment
func FormattedValues(){
	//Printf does not add new line like Println
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("An string: %s\n", "hello")
	fmt.Printf("An boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2,"\t",true)
	//Prints values as go sees them
	//Useful for debug
	//\t is a tab
	fmt.Printf("Values: %#v %#v %#v\n", 1.2,"\t",true)
	fmt.Printf("Type: %T %T %T\n", 1.2,"\t",true)
	fmt.Printf("Percent sign: %%\n")

	fmt.Println("--------------------------")

	//Can use char widths with Printf to set character lengths
	//example string 12 wide or integer 2 wide
	//and width for decimal points
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------------------")
	fmt.Printf("%12s | %2d\n","Stamps", 50)
	fmt.Printf("%12s | %2d\n","Paper clips", 5)
	fmt.Printf("%12s | %2d\n","Tape", 99)

	fmt.Println("--------------------------")

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	//Can use this to round to two decimal places
	fmt.Printf("%%.2f: %.2f\n", 12.3456)

}