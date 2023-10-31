package main

import "fmt"

func main() {
	var o1, o2, o3 string
	s1 := "Кот"
	s2 := "Жираф"
	s3 := "Курица"
	s4 := "Страус"
	s5 := "Дельфин"
	s6 := "Плезиозавры"
	s7 := "Пингвин"
	s8 := "Утка"
	fmt.Scan(&o1, &o2, &o3)
	if o1 == "Нет" && o2 == "Нет" && o3 == "Нет" {
		fmt.Println(s1)
	}
	if o1 == "Нет" && o2 == "Нет" && o3 == "Да" {
		fmt.Println(s2)
	}
	if o1 == "Нет" && o2 == "Да" && o3 == "Нет" {
		fmt.Println(s3)
	}
	if o1 == "Нет" && o2 == "Да" && o3 == "Да" {
		fmt.Println(s4)
	}
	if o1 == "Да" && o2 == "Нет" && o3 == "Нет" {
		fmt.Println(s5)
	}
	if o1 == "Да" && o2 == "Нет" && o3 == "Да" {
		fmt.Println(s6)
	}
	if o1 == "Да" && o2 == "Да" && o3 == "Нет" {
		fmt.Println(s7)
	}
	if o1 == "Да" && o2 == "Да" && o3 == "Да" {
		fmt.Println(s8)
	}
}
