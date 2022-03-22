package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("b no puede ser 0")
	}
	return a / b, nil

}

func main() {
	div, err := division(7, 0)
	if err != nil {
		fmt.Println("Error en la division", err.Error())
	}
	fmt.Println("division", div)
}
