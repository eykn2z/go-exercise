package main

import "fmt"

/*
===
prototype
===
正確にオブジェクトをコピーするオブジェクト

オブジェクトのメンバーの一部がinterfaceだと、違うinterfaceに差し替えて作成も可能になってしまう
privateメンバーの存在等

ほとんどの言語では同じクラスのオブジェクトは他のオブジェクトの非公開フィールドにもアクセス可能

自分のコードが外部開発コードからのオブジェクトと動作する場合によく使う・依存させたくない
オブジェクトの初期化方法のみが異なるサブクラスの数を減らしたい場合に使う

ts: Object.assign
python: copy.copy
で利用可能

https://refactoring.guru/ja/design-patterns/prototype/go/example
*/

type Inode interface {
	print(string)
	clone() Inode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{folder1, file2},
		name:     "Folder2",
	}
	cloneFolder := folder2.clone()
	cloneFolder.print("  ")
}
