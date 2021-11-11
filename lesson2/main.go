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

	// byte
	by := []byte{74, 72}
	fmt.Println(by)
	fmt.Println(string(by))

	// 配列
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{77, 88, 99}
	fmt.Println(arr3)

	// ...で可変長にできる
	arr4 := [...]string{"Go", "Lang"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// interface型
	// あらゆる型と互換性がある
	var x interface{}
	fmt.Println(x) // <nil>が表示される 何も値がないよの意味
	fmt.Printf("%T\n", x)

	x = 100
	fmt.Println(x)
	fmt.Printf("数値 %T\n", x)

	x = "XYZ"
	fmt.Println(x)
	fmt.Printf("文字列 %T\n", x)

	x = [3]int{1, 3, 6}
	fmt.Println(x)
	fmt.Printf("配列 %T\n", x)

	// interface型は様々な値をいれることができるがデータ特有の計算などができない
	// 以下はエラーになる
	// x = 8
	// fmt.Println(x + 3)

}
