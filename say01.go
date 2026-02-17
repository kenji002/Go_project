package main

import (
	"fmt"
	"reflect"

	"example.com/myapp/utils"
)

// ターミナルで「go run say01.go」で実行
// フォルダ内では「go run .」で実行可能
func main() {
	Function_1()
	fmt.Println(reflect.TypeOf(Function_1))

	Function_2()
	fmt.Println(reflect.TypeOf(Function_2))

	fmt.Println(utils.Add(1, 2))
	utils.PrintAdd(1, 2)
}

func Function_1() {
	var integer int = 1 + 1
	var str string = "こんにちは" + " さようなら"
	var floating float64 = 3.14 + 3.14
	var boolean bool = true && false

	fmt.Println(integer)
	fmt.Println(str)
	fmt.Println(floating)
	fmt.Println(boolean)
}

func Function_2() {
	integer := 1 + 1
	str := "こんにちは" + " さようなら"
	floating := 3.14 + 3.14
	boolean := true && false

	fmt.Println(integer)
	fmt.Println(str)
	fmt.Println(floating)
	fmt.Println(boolean)
}
