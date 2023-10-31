package main

import "fmt"

func main() {
	var x, i int
	fmt.Scan(&x)
	for ; x != 1; i++ {
		if x%2 == 0 {
			x = x / 2
		} else if x%2 != 0 {
			x = 3*x + 1
		}
	}
	fmt.Println(i)
}
