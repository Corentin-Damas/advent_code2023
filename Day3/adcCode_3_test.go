package day3

import (
	"fmt"
	"testing"
)

func TestLineToarray(t *testing.T) {
	stringTotest := "467..114.."
	
	got, _ := lineToarray(stringTotest)

	runeTotest := []rune{}
	strToRune := []rune("467..114..")

	for _, char := range strToRune {
		runeTotest = append(runeTotest, char)
	}
	want := runeTotest

	// err := errors.New("char don't match" )

	for i := 0; i <= len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted %q", got[i], want[i])

		}
	}
}


func TestSymboleLoc(t * testing.T){
	var arr = []rune{54, 49, 55, 42, 46, 46,46, 46, 46, 46}

	got := symboleLoc(arr)

	var want = []int{3} 

	if len(got) != len(want){
		t.Errorf("got %q, wanted %q", got, want)
	}
	fmt.Println(got)

}


func TestCheckFullNumber(t * testing.T){
	var arr = []rune{54, 49, 55, 42, 46, 46,46, 46, 46, 46}

	got := checkFullNumber(arr, 3)

	want := 617

	if got != want{
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDay3result(t * testing.T){
	got := Day3result()
	want := 540887

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}