package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	//this is an anonymous field
	//it is promoted to outer struct
	//however in this case date fields are not promoted as they are encapsulated = lowercase (not accessible)
	//cannot do event.month or event.Date.month
	//however the exported methods do get promoted to the outer struct
	//so can do
	//date := calendar.Event{}
	//date.SetYear(1977)
	//date.Year()
	Date
}

func (e *Event) SetTitle(title string) error{
	//checks numbers of chars in title
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string{
	return e.title
}