package main

import "fmt"

/*
===
bridge
===
巨大なクラスや密接に関連したクラスの集まりを抽象部分と実装部分という２つの階層に分離
レイヤードアーキテクチャに使う交換可能なコンポーネントの組み方と同様？
Bridgeはアプリケーションの部分部ぬんを独立して開発できるように、設計当初から使われる
Adapterは既存のアプリケーションに対して利用され、互換性のないクラスとうまく動作させるために使う
Bridge, State,Strategyはよく似ているが違う問題解決をする

抽象化・実装どの組み合わせでも動作する
■の青、赤、●の青、赤
*/

func main() {
	blue := &Blue{}
	red := &Red{}

	square := &Square{}
	circle := &Circle{}

	square.SetColor(blue)
	square.Print()

	circle.SetColor(red)
	circle.Print()

	// この場合SetColorを指定しないとエラーになる
	// 実際はconstructorに入れた方がいいが、抽象体を渡す
	square2 := &Square{}
	square2.Print()
}

type Shape interface {
	Print()
	SetColor(Color)
}

type Square struct {
	color Color
}

func (s *Square) Print() {
	fmt.Println("Print for Square")
	s.color.Print()
}

func (s *Square) SetColor(c Color) {
	s.color = c
}

type Circle struct {
	color Color
}

func (s *Circle) Print() {
	fmt.Println("Print for Circle")
	s.color.Print()
}

func (s *Circle) SetColor(c Color) {
	s.color = c
}

type Color interface {
	Print()
}

type Blue struct{}

func (b *Blue) Print() {
	fmt.Println("Print for Blue")
}

type Red struct{}

func (b *Red) Print() {
	fmt.Println("Print for Red")
}
