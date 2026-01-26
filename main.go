package main

import "fmt"

type Vehicle struct {
	Speed int
}

type Bus struct {
	Vehicle
}

func (v *Vehicle) Ride() {
	fmt.Printf("I am default behavior\n")
}

// func (b *Bus) Ride() {
// 	fmt.Printf("I am a bus! I have no speed:(\n")
// }

func main() {
	v := &Vehicle{
		Speed: 100,
	}
	b := Bus{}

	v.Ride()
	b.Ride()
}
