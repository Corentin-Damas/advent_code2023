package day3

import (
	"testing"
)

func TestLineToarray(t *testing.T) {
	stringTotest := "467..114.."

	got := lineToarray(stringTotest)

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
	var arr = []rune{46,46,46,36,46,42,46,46,46,46}

	got := symboleLoc(arr)

	var want = []int{3, 5} 

	if len(got) != len(want){
		t.Errorf("got %q, wanted %q", got, want)
	}

}
