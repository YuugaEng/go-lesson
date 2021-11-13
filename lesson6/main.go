package main

import "fmt"

// 関数
// ジェネレーターの実装

func main() {
	ints := integers()
	fmt.Println("ints")
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// 再定義した場合は値もクリアされる
	otherints := integers()
	fmt.Println("otherints")
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
