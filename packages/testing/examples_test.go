package testing

import (
	"testing"
)

//tests the expected output from JoinWithCommas with two elements
func TestTwoElements(t *testing.T){
	list := []string{"apple","orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list,got,want)
	}
}

//tests the expected output from JoinWithCommas with three elements
func TestThreeElements(t *testing.T){
	list := []string{"apple","orange","pear"}
	want := "apple, orange and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list,got,want)
	}
}