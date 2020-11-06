package testing

import (
	"fmt"
	"testing"
)

//struct to hold data for tests
//contains input and output for each test
type testData struct{
	list []string
	want string
}

//you can use a struct to hold the inout and output of each test
func TestJoinWithCommas(t *testing.T){
	//create a string of testData
	tests := []testData{
		//you can omit testData as go know the type from parent
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple","orange"}, want: "apple and orange"},
		{list: []string{"apple","orange","pear"}, want: "apple, orange and pear"},
	}
	//for each test
	for _,test := range tests{
		//get the string from the list
		got := JoinWithCommas(test.list)
		//test if what you got is what you want
		if got != test.want {
			t.Errorf(errorString(test.list,got,test.want))
		}
	}
}

/*
func TestOneElement(t *testing.T){
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list,got,want))
	}
}

//tests the expected output from JoinWithCommas with two elements
func TestTwoElements(t *testing.T){
	list := []string{"apple","orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list,got,want))
	}
}

//tests the expected output from JoinWithCommas with three elements
func TestThreeElements(t *testing.T){
	list := []string{"apple","orange","pear"}
	want := "apple, orange and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list,got,want))
	}
}
*/


//helper functions can be created to help refactor code
//example this one format error string used across multiple tests
func errorString(list []string, got string, want string) string{
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list,got,want)
}