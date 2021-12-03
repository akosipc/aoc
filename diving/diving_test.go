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

	if ans != 900 {
		t.Errorf("divingPos() = %d; want 900", ans)
	}
}

func BenchmarkDivingBasic(b *testing.B) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	for i := 0; i < b.N; i++ {
		divingPos(input)
	}
}
