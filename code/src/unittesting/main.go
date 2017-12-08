package main

import "fmt"

func Times(x int, y int) int {
	return x * y
}

func main() {
	t := Times(5, 5)
	fmt.Printf("%d", t)
}
