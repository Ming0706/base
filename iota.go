package main

import "fmt"

func main() {
	const (
		a = iota; b
		c = 3
		d
	)
	fmt.Printf("a=%v, b=%v, c=%v, d=%v \n", a, b, c, d)
}
