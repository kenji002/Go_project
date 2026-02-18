package utils

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintAdd(a int, b int) {
	fmt.Println(Add(a, b))
}

func Sub(a, b int) int {
	return a - b
}

func PrintSub(a int, b int) {
	fmt.Println(Sub(a, b))
}
