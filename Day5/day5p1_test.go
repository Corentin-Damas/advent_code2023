package day5

import (
	"fmt"
	"testing"
)

func TestGetSeeds(t *testing.T) {
	got := getSeeds("seeds: 79 14 55 13")
	want := []int{79, 14, 55, 13}

	fmt.Println(got, "want : ", want)

	for i := 0; i <= len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("got %d, wanted %d", got[i], want[i])

		}

	}
}

func TestGetStepTitle(t *testing.T) {
	got := getStepTitle("soil-to-fertilizer map:")
	want := "fertilizer"
	if got != want {
		t.Errorf("got %s,  wanted %s", got, want)
	}
}

func TestUnmap(t *testing.T) {
	got := unMap(98, 5, -48)
	want := []int{98, 102, -48}

	for i := 0; i <= len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("got %d, wanted %d", got[i], want[i])
		}
	}
}
