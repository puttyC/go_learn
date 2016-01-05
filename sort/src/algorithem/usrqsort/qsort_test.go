package usrqsort

import "testing"

func TestQsort_GivenInvertOrderWhenQsortThenOrder(t *testing.T) {
	values := []int{1, 3, 2, 5, 4}
	Qsort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("Qsort failed ! Got values", values, " Expected 1 2 3 4 5")
	}
}
func TestQsort_GivenOrderWhenQsortThenOrder(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	Qsort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("Qsort failed ! Got values", values, " Expected 1 2 3 4 5")
	}
}

func TestQsort_GivenInvertOrderWithSameElemWhenQsortThenOrder(t *testing.T) {
	values := []int{1, 1, 2, 5, 2}
	Qsort(values)
	if values[0] != 1 || values[1] != 1 || values[2] != 2 || values[3] != 2 || values[4] != 5 {
		t.Error("Qsort failed ! Got values", values, " Expected 1 1 2 2 5")
	}
}
func TestQsort_GivenOneElemWhenQsortThenOrder(t *testing.T) {
	values := []int{5}
	Qsort(values)
	if values[0] != 5 {
		t.Error("Qsort failed ! Got values", values, " Expected 5")
	}
}
