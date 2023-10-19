package main

import "log"

/*
自分が書いたライブラリやコンポーネントに対して、ユーザーがカスタムコンポーネントを指定出来るようにしたい場合
インスタンスを生成する場合はファクトリーメソッドで置き換える
テスト時にテスト用のinstanceに差し替えが可能
*/

type IAnimal interface {
	setName(name string)
	getName() string
	setKind(kind string)
	getKind() string
}

type Cat struct {
	name string
	kind string
}

func (c *Cat) setName(name string) {
	c.name = name
}

func (c *Cat) getName() string {
	return c.name
}

func (c *Cat) setKind(kind string) {
	c.kind = kind
}

func (c *Cat) getKind() string {
	return c.kind
}

func newAnimal(name, kind string) IAnimal {
	// Animal生成をTigerにしたい場合、このメソッドのみを変更すればよい
	return &Cat{name: name, kind: kind}
}

func main() {
	cat := newAnimal("Siro", "Abyssinian")
	log.Printf("%d", cat)
}
