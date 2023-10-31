package main

import "fmt"

func main() {
	var k int
	var a, b, c float64
	fmt.Scan(&k)
	var m []int = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&m[i])
	}
	for j := 0; j < k; j++ {
		if j == 0 {
			a = float64(m[j]) * 1.0
			fmt.Print(float64(a), " ")
		}
		if j > 0 && j < k-1 {
			b = (float64(m[j]) + float64(m[j+1]) + float64(m[j-1])) / 3.0
			fmt.Print(b, " ")
		}
		if j == k-1 {
			c = float64(m[j]) * 1.0
			fmt.Print(c, " ")
		}
	}
}
