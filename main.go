package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,World!123")
	fmt.Println("lixiong")
	fmt.Println("Day0416")
	subab(2, 3)
	sumab(2, 4)
	fmt.Println("123")
}

func sumab(a, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

func subab(a, b int) {
	sub := a - b
	fmt.Println(sub)
}
