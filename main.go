package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10

	fmt.Println("Woow!")
	fmt.Println("So cool")

	fmt.Println(reflect.TypeOf(a))
}
