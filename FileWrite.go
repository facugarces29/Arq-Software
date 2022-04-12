package main

import (
	"fmt"
	"os"
)

func main() {

	path := "C:\\Users\\facuu\\Desktop\\Arq-Software\\algo.txt"
	f, _ := os.Create(path) //*os.file
	fmt.Fprintf(f, "data")
	f.Close()

}
