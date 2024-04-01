package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var number int64 = 6
	var bitInx int = 1

	number |= (1 << bitInx)
	fmt.Printf("После установки %d-го бита в 1 число равно %d\n", bitInx+1, number)
	printBinary(number)

	fmt.Println("_____________________________")

	number &= ^(1 << bitInx)
	fmt.Printf("После установки %d-го бита в 0 число равно %d\n", bitInx+1, number)
	printBinary(number)
}

func printBinary(number int64) {
	res, err := Convert(strconv.Itoa(int(number)), 10, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Двоичное представление числа: ", res)
}

func Convert(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
