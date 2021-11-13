package main

import "fmt"

// 関数
// クロージャーの実装

func main() {
	f := Later()
	fmt.Println(f("Hello"))  // 一番最初に呼ばれたのでLater関数で定義しているstoreの初期値が入る
	fmt.Println(f("My"))     // ↑で呼び出した際にHelloを渡しているのでHelloが帰ってくる
	fmt.Println(f("name"))   // My
	fmt.Println(f("is"))     // name
	fmt.Println(f("Golang")) // is
}

// 前回のデータを保持して返す
//
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
