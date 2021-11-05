package main

import "testing"

func TestPlay(t *testing.T) {
	input := 389125467

	got := play(intToIntArray(input))
	gotInt := intArrayToInt(got)

	target := []int{2, 3, 5, 7, 11, 13}
	gotTarget := intArrayToInt(target)

	if gotInt != gotTarget {
		t.Errorf("got: %d wanted: %d", got, target)
	}
}
