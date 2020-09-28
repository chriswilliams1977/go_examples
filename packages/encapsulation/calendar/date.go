package calendar

import(
	"errors"
)

type Date struct {
	year int
	month int
	day int
}

//Setters
//Make sure people can only update through getters and setters by making type fields unexported
//thus they cannot be referenced directly
//this is ENCAPSULATION = hiding one part of code from another part
//when  calling this you dont need to pass pointer, go automatically does this for you
//if reciever is a pointer
func (d *Date) SetYear(year int) error{
	//check fortmat of year
	if year <1 {
		//return error if invalid year
		return errors.New("Invalid year ")
	}
	//go automatically get pointer value, yopu dont need to do (*d).Year
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error{
	if month <1 || month >12{
		return errors.New("Invalid month ")
	}
	//go automatically get pointer value, yopu dont need to do (*d).Year
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error{
	if day <1 || day >31{
		return errors.New("Invalid day ")
	}
	//go automatically get pointer value, yopu dont need to do (*d).Year
	d.day = day
	return nil
}

//Getters
//For getters the name is the same as the field you are getting
//If you pointers for setters you should do the same for getters
func (d *Date) Year () int{
	return d.year
}

func (d *Date) Month () int{
	return d.month
}

func (d *Date) Day () int{
	return d.day
}