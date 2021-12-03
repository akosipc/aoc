package main

import (
	"testing"
)

func TestCalculateIncreaseBasic(t *testing.T) {
	slice := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	ans := determineIncrease(slice)

	if ans != 7 {
		t.Errorf("calculateIncreaseBasic() = %d; want 7", ans)
	}
}

func TestThreeMeasurementWindowSlice(t *testing.T) {
	slice := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	ans := threeMeasurementWindowSlice(slice)

	if ans != 5 {
		t.Errorf("threeMeasurementWindowSlice = %d; want 5", ans)
	}
}

func BenchmarkCalculateIncrease(b *testing.B) {
	slice := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	for i := 0; i < b.N; i++ {
		determineIncrease(slice)
	}
}
