package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is New World")

	for index, i := range "Naveen Hiremath" {
		fmt.Printf("%d -> %c\n", index, i)
	}
	timee := displayDate()
	fmt.Println(timee)
}
