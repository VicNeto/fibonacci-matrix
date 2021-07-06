package service

import (
	"reflect"
	"testing"
)

func FibonnaciSpiralTest(t *testing.T) {
	cases := []struct {
		rows   int
		cols   int
		spiral [][]int
	}{
		{1, 1, [][]int{{1}}},
		{1, 2, [][]int{{1, 1}}},
		{2, 2, [][]int{{1, 1}, {2, 3}}},
		{3, 3, [][]int{{1, 1, 2}, {3, 5, 8}, {13, 21, 34}}},
	}

	for _, c := range cases {
		a, err := FibonnaciSpiral(c.rows, c.cols)
		if err != nil {
			// t.Fatalf(“error should be nil”)
			t.Fail()
		}
		if !reflect.DeepEqual(a, c.spiral) {
			// t.Log(“error should be”+expected+”, but got”,str)
			t.Fail()
		}
	}
}
