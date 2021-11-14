package main

import "fmt"

func main() {

	// switch

	// 列挙
	n := 5
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don't know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don't know")
	}

	// 演算子
	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n2 < 4")
	case n2 > 3 && n2 < 7:
		fmt.Println("3 < n2 < 7")
	}

	/*
		// caseに列挙と演算子を混同させた場合はエラーになる
		switch n := 2; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		case n > 3 && n < 6:
			fmt.Println("3 < n < 6")
		default:
			fmt.Println("i don't know")
		}
	*/

}
