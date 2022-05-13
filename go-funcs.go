package main

import "fmt"

func sayHi() {
	fmt.Println("HI")
}

func wish() {
	fmt.Println("Good Day!")
	var a int
	// fmt.Scanf("%d", &a)
	for i := 0; i < 3000000; i++ {
		a++
		if a%2000 == 0 {
			fmt.Println("Good Day!", a)
		}
	}
}

func main() {
	// fmt.Println("in start")

	go sayHi()
	wish()
}
