package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{}

	value = 42
	checkType(value)
	checkTypeWithReflect(value)

	value = "Hello, world!"
	checkType(value)
	checkTypeWithReflect(value)

	value = true
	checkType(value)
	checkTypeWithReflect(value)

	ch := make(chan int)
	value = ch
	checkType(value)
	checkTypeWithReflect(value)
}

func checkType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("Это переменная типа int | typeswitch")
	case string:
		fmt.Println("Это переменная типа string | typeswitch")
	case bool:
		fmt.Println("Это переменная типа bool | typeswitch")
	case chan int:
		fmt.Println("Это переменная типа channel | typeswitch")
	default:
		fmt.Println("unknow type")
	}
}

func checkTypeWithReflect(i interface{}) {
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("Это переменная типа int | reflect\n\n")
	case reflect.String:
		fmt.Printf("Это переменная типа string | reflect\n\n")
	case reflect.Bool:
		fmt.Printf("Это переменная типа bool | reflect\n\n")
	case reflect.Chan:
		fmt.Printf("Это переменная типа channel | reflect\n\n")
	default:
		fmt.Println("Неизвестный тип:", t)
	}
}
