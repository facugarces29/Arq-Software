package main

import (
	"os"
	"fmt"
)

func main() {
	path := "C:\\Users\\facuu\\Desktop\\Arq-Software\\algo1.txt"
	f, _ := os.ReadFile(path)
	fmt.Printf(string(f))
	f.Close(f

}
