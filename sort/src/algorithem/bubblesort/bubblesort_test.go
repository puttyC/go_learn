package bubblesort

import "testing"

func TestBubbleSort_GivenValuesSizeZero(t *testing.T) {
	values := []int{}
	result := BubbleSort(values)
	if result != 0 {
		t.Error("BubbleSort failed! Given values size is zero! Expected 0")
	}
}

func TestBubbleSort_GivenValuesNormal(t *testing.T) {
	values := []int{1, 4, 3, 2, 5}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort failed! Given values ", values, "! Expected 1,2,3,4,5")
	}
}

func TestBubbleSort_GivenValuesHaveSameElement(t *testing.T) {
	values := []int{1, 5, 3, 5, 5}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 3 || values[2] != 5 || values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort failed! Given values ", values, "! Expected 1,3,5,5,5")
	}
}

func TestBubbleSort_GivenValuesOrder(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort failed! Given values ", values, "! Expected 1,2,3,4,5")
	}
}

func TestBubbleSort_GivenValuesInvertOrder(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort failed! Given values ", values, "! Expected 1,2,3,4,5")
	}
}
func TestBubbleSort_GivenValuesSizeOne(t *testing.T) {
	values := []int{1}
	BubbleSort(values)
	if values[0] != 1 {
		t.Error("BubbleSort failed! Given values ", values, "! Expected 1")
	}
}
