package main

import (
	"fmt"
	"strconv"
)

// 型変換のレッスン
func main() {

	// 数値
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	fl64 = 5.2
	i2 := int(fl64)
	fmt.Println(i2)
	fmt.Printf("i2 = %T\n", i2)

	// 文字
	var s string = "100"
	fmt.Println(s)
	fmt.Printf("s = %T\n", s)

	// 文字はstrconv.Atoi(変数)で数値へ変換できる
	// 注意として戻りは2つある
	// 数値に変換された値とエラーがある
	// エラーの値は使わないため、「_」を設定することで捨てている
	i3, _ := strconv.Atoi(s)
	fmt.Println(i3)
	fmt.Printf("i3 = %T\n", i3)

	var i4 int = 200
	s2 := strconv.Itoa(i4)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

}
