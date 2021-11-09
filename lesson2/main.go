package main

import "fmt"

func main() {

	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go!"
	fmt.Println(s)

	// 複数定義
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 10000
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 値の設定なし
	// int型は0が入る string型は空文字が入る bool型はfalseが入る
	var i3 int
	var s3 string
	var b bool
	fmt.Println(i3, s3, b)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 500
	fmt.Println(i4)

	// 値の代入
	i4 = 550
	fmt.Println(i4)

	// 一番最初に設定した値の型が自動で入る
	// 今回は500で数値で設定したためint型になっている
	// そのため下のように文字の値の設定や暗黙的な定義の再定義はエラーになる
	// i4 = "Error"
	// i4 := "Error"
	// 関数外での暗黙的な定義はできない
}
