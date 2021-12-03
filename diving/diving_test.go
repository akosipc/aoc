package main

import "testing"

func TestStringToIntBasic(t *testing.T) {
	str := "7"

	ans := stringToInt(str)

	if ans != 7 {
		t.Errorf("stringToInt() = %d; want 7", ans)
	}
}

func TestDivingBasic(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	ans := divingPos(input)

	if ans != 150 {
		t.Errorf("divingPos() = %d; want 900", ans)
	}
}
