package day9

import (
	"fmt"
	"testing"
)

func TestToNumsArr(t *testing.T) {
	got := toNumsArr("1 3 6 10 15 21")
	want := []int{1, 3, 6, 10, 15, 21}

	fmt.Println(got, "want : ", want)

	for i := 0; i <= len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("got %d, wanted %d", got[i], want[i])

		}

	}
}