package day2

import (
	"testing"
)

func TestGameId(t * testing.T){
	got := gameId("Game 12")
	want := 12
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestColorCheck(t * testing.T){
	bag := Bag{
		red: 12,
		green: 13,
		blue: 14,
	}
	var pBag *Bag = &bag

	err := colorCheck("13 green", pBag)
	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

}

