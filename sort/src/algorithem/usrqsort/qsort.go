package usrqsort

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	tmp := a[left]
	i, j := left, right
	/*go 仅支持for，不支持while和do-while*/
	for {
		for i < j && a[j] >= tmp {
			j--
		}
		if i < j {
			a[i] = a[j]
		} else {
			break
		}
		for i < j && a[i] <= tmp {
			i++
		}
		if i < j {
			a[j] = a[i]
		} else {
			break
		}
	}
	a[i] = tmp
	quickSort(a, left, i-1)
	quickSort(a, i+1, right)
}

func Qsort(a []int) {
	quickSort(a, 0, len(a)-1)
}
