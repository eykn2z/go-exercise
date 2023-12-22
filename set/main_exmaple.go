package set

import (
	"fmt"
)

func main() {
	// 新しいsetを作成
	mySet := New()

	// 要素を追加
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("orange")

	// 要素が含まれているか確認
	fmt.Println("Contains 'apple':", mySet.Contains("apple")) // trueが返る
	fmt.Println("Contains 'grape':", mySet.Contains("grape")) // falseが返る

	// setの要素を表示
	fmt.Println("Set items:", mySet.Items())

	// 要素を削除
	mySet.Remove("banana")
	fmt.Println("Contains 'banana' after removal:", mySet.Contains("banana")) // falseが返る
	fmt.Println("Set items after removal:", mySet.Items())
}
