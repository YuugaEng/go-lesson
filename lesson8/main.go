package main

import "fmt"

func main() {

	// for文

	// 無限ループ
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("loop")
	}

	i2 := 0
	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// 配列を使った応用
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// indexは使わない場合
	// 値のみほしい
	arr2 := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr2 {
		fmt.Println(v)
	}

	// Mapを使った応用
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// valueのみほしい
	for _, v := range m {
		fmt.Println(v)
	}

	// keyのみほしい
	// for k, _ を単純化
	for k := range m {
		fmt.Println(k)
	}

}
