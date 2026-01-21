package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	a := "sfsf"
	m := map[int]int{1: 10, 2: 20, 3: 30}

	fmt.Println(s)
	modifySlice(s)
	fmt.Println(s)

	println("")

	fmt.Println(a)
	modifyStr(a)
	fmt.Println(a)

	println("")

	fmt.Println(m)
	modifyMap(m)
	fmt.Println(m)
}

func modifySlice(s []int) {
	s[0] = 100
}

/*
 * type Slice struct {
 * 	pointer *[n]int
 * len
 * cap
 * }
 */

func modifyStr(a string) {
	a = "sfafsfdsgsgsdfsdfs"
	fmt.Println(a)
}

func modifyMap(m map[int]int) {
	m[1] = 12345
}
