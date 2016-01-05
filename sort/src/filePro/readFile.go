package filePro

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadValue(file *string) (values []int, err error) {
	infile, err := os.Open(*file)
	defer infile.Close()
	if err != nil {
		fmt.Println("Can't open infile=", *file, " ,erro= ", err)
	}
	readLine := bufio.NewReader(infile)
	values = make([]int, 0)
	for {
		line, isPrefix, err := readLine.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read file occur error, err=", err)
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seem unexpected,isPrefix = ", isPrefix)
		}
		str := string(line)
		value, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("value process occur err, err=", err)
			break
		}
		values = append(values, value)
	}
	return
}
