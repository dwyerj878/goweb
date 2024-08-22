package main

import "testing"

func TestAdd(t *testing.T) {
	x := 1
	y := 10
	if add(x, y) != 11 {
		t.Error("did not add")
	}

}
