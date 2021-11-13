package main

import "fmt"

func main() {

	f := ReturnFunc()
	f()

	CallFunction(func() {
		fmt.Println("I'm a call function")
	})

}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}
