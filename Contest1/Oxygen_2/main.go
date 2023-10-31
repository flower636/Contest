package main

import "fmt"

func main() {
	h := 0.5 * 365
	kt := h/32 + 1
	kd := h/20 + 1
	fmt.Println(h, int(kt), int(kd))
}
