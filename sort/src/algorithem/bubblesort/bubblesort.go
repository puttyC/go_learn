package bubblesort

import "fmt"

func BubbleSort(values []int) (resultFlag int) {
	if len(values) == 0 {
		fmt.Println("sorting input values size is zero")
		resultFlag = 0
		return
	}
	for i := len(values) - 1; i >= 0; i-- {
		var flg bool = true
		for j := 0; j < i; j++ {
			if values[j] > values[j+1] {
				values[j+1], values[j] = values[j], values[j+1]
				flg = false
			}
		}
		if flg {
			break
		}
	}
	resultFlag = 0
	return
}
