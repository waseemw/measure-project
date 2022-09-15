package main

import "fmt"

func slices() {
	b := []int{1, 2, 3}
	c := b[1:2]

	b[1] = 3
	fmt.Println(b)
	fmt.Println(c)
}
