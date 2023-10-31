package main

import "fmt"

func main() {
	var row, col int
	var t4 string
	t2 := "--"
	t4_start := "----"
	fmt.Scan(&row, &col)
	//Вывод первой строки
	if row >= 1 && row <= 31 && col >= 1 && col <= 31 {
		fmt.Print("    |")
		for k := 1; k <= col; k++ {
			if k < 10 {
				fmt.Print("   ")
				fmt.Print(k)
			} else {
				fmt.Print("  ")
				fmt.Print(k)
			}
		}
		fmt.Println("")
		//Вывод строки минусов
		fmt.Print("   ")
		for c := 0; c < col; c++ {
			t4 = t4 + t4_start
		}
		fmt.Print(t2)
		fmt.Println(t4)
		//вывод остальной таблицы
		for i := 1; i <= row; i++ {
			if i < 10 {
				fmt.Print("   ", i, "|")
				for j := 1; j <= col; j++ {
					if i*j < 10 {
						fmt.Print("   ")
						fmt.Print(i * j)
					} else if i*j >= 10 && i*j < 100 {
						fmt.Print("  ")
						fmt.Print(i * j)
					} else {
						fmt.Print(" ")
						fmt.Print(i * j)
					}
				}
			} else {
				//fmt.Print(" ")
				fmt.Print("  ", i, "|")
				for j := 1; j <= col; j++ {
					if i*j < 10 {
						fmt.Print("   ")
						fmt.Print(i * j)
					} else if i*j >= 10 && i*j < 100 {
						fmt.Print("  ")
						fmt.Print(i * j)
					} else {
						fmt.Print(" ")
						fmt.Print(i * j)
					}
				}
			}
			fmt.Println("")
		}
	}
}
