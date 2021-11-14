package main

import (
	"fmt"
	"strconv"
)

func main() {

	// if文
	// ()は省略できる

	a := 0

	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 5; true {
		// ここで定義されているxとifの外で定義されているxは別物
		fmt.Println(x)
	}
	// ifで5を代入されているがそれはif文内だけのため0のまま
	fmt.Println(x)

	var s string = "100"
	i, err := strconv.Atoi(s)
	// 変換に成功した場合はerrにはnilが入る
	// 変換に失敗した場合はerr内容が入る
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	var s2 string = "ABC"
	i2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i2 = %T\n", i2)

}
