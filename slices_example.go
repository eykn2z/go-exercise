package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	// スライスの初期化
	numbers := []int{1, 2, 3, 4, 5}
	strings := []string{"hello", "world"}
	fmt.Println("Original slice:", numbers)

	// スライスの逆順
	slices.Reverse(numbers)
	fmt.Println("Reversed slice:", numbers)
	slices.Reverse(strings)
	fmt.Println("Reversed slice:", strings)

	// 特定の要素を含むかチェック
	containsThree := slices.Contains(numbers, 3)
	fmt.Println("Slice contains 3:", containsThree)

	// 特定の要素のインデックスを探す
	indexOfThree := slices.Index(numbers, 3)
	fmt.Println("Index of 3:", indexOfThree)

	// スライスのコピー
	copiedNumbers := slices.Clone(numbers)
	fmt.Println("Copied slice:", copiedNumbers)
}
