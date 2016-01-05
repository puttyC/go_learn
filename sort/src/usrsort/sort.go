package main

import (
	"algorithem/bubblesort"
	"algorithem/usrqsort"
	"filePro"
	"flag"
	"fmt"
	"reflect"
)

var infile *string = flag.String("i", "unsort.dat", "File contain values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")

var algorithem *string = flag.String("a", "qsort", "Sort algorithem")

func main() {
	/*1.参数解析*/
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile,
			"algorithem = ", *algorithem)
	}
	/*2.读入待排序的数据*/
	var values []int
	var err error
	values, err = filePro.ReadValue(infile)
	if err != nil {
		fmt.Println("read File failed!")
	}
	fmt.Println(values)
	var a int = 5
	fmt.Println("int size = ", reflect.TypeOf(a).Size())
	/*3.根据算法参数进行排序*/
	switch *algorithem {
	case "qsort":
		usrqsort.Qsort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("Sort algorithem ", *algorithem, "is either unknow or unsupported")
	}
	fmt.Println(values)
	/*4.将结果写入到输出文件中*/
	err = filePro.WriteFile(values, outfile)
}
