package interfaces

import "fmt"

type OverheatError float64

//this implements the error interface
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

//NOTE
//error is a predeclared identifier
//means its is available universally - you dont need to import a package
func CheckTemperature(actual float64, safe float64) error {
	//check execess temp
	excess := actual - safe
	//if over 0 then return OverheatError
	if excess > 0 {
		//you can return this because OverheatError implements error interface
		//OverheatError(excess) is the OverheatError type set with value (float64)
		return OverheatError(excess)
	}
	return nil
}
