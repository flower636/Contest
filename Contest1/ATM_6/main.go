package main

import "fmt"

func main() {
	var n, k5000, k1000, k500, k200, k100 int
	fmt.Scan(&n)
	for i := 0; n > 0; i++ {
		if n >= 5000 {
			n = n - 5000
			k5000++
		} else if n >= 1000 {
			n = n - 1000
			k1000++
		} else if n >= 500 {
			n = n - 500
			k500++
		} else if n >= 200 {
			n = n - 200
			k200++
		} else if n >= 100 {
			n = n - 100
			k100++
		}
	}
	fmt.Println(k5000, k1000, k500, k200, k100)
}
