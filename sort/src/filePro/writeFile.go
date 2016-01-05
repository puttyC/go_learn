package filePro

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFile(values []int, file *string) (err error) {
	outfile, err := os.Create(*file)
	defer outfile.Close()
	if err != nil {
		fmt.Println("Can't creat outfile = ", *file, " err=", err)
		return err
	}
	for _, value := range values {
		str := strconv.Itoa(value)
		size, err := outfile.WriteString(str + "\n")
		if err != nil {
			fmt.Println("write File occur error, error=", err)
			return err
		}
		fmt.Println("already write size = ", size)
	}
	return nil
}
