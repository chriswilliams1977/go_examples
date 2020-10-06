package interfaces

import "fmt"

type Gallons float64
//implements the Stringer interface which has String() method which returns a string
//Stringer interface lets types define how they will be displayed when printed
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
