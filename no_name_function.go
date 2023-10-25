package main

import "fmt"

//無名関数
func main() {
	func(str string) {
		fmt.Printf(str)
	}("test")

	func() {
		fmt.Printf("test 2")
	}()
}
