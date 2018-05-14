package main

import "fmt"

func Times(x int, y int) int {
	return x * y
}

func main() {
	t := Times(5, 5)
	o := Devide(10, 5)
	fmt.Printf("%d", t)
	fmt.Printf("%d", o)
}
