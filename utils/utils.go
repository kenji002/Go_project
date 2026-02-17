package utils

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintAdd(a int, b int) {
	fmt.Println(Add(a, b))
}
