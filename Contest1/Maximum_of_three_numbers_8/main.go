package main

import "fmt"

func main() {
	var a, b, c, max int
	fmt.Scan(&a, &b, &c)
	max = 0
	if a > max {
		max = a
	}
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	fmt.Println(max)
}
