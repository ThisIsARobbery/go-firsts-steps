package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Ваше имя: ")
	fmt.Scan(&name)
	fmt.Print("Ваш возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}
