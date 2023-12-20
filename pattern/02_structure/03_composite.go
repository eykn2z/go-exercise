package main

import "fmt"

/*
===
composite
===
ツリー構造の全ての内容物に対して再起的にある振る舞いを実行可能
ループを使わなくて済む
ツリー構造の時に実装する
構成要素
- 単純項目
- コンテナ
- 単純要素と複合要素の両方に意味のあるメソッドを並べたコンポーネント・インターフェース

開放閉鎖の原則　既存のコードを壊すことなく、オブジェクトツリーと動作可能な新規の要素型をアプリに導入可能
*/

func main() {
	person := &Person{name: "John"}
	person2 := &Person{name: "Watson"}
	person3 := &Person{name: "Hornet"}

	businessDepartment := &Department{
		name: "Business",
	}
	businessDepartment.add(person)
	developDepartment := &Department{
		name: "Develop",
	}
	businessDepartment.add(developDepartment)
	developDepartment.add(person2)
	developDepartment.add(person3)

	res := businessDepartment.search("Rose")
	fmt.Printf("found: %s\n", res)
	res = businessDepartment.search("Hornet")
	fmt.Printf("found: %s\n", res)

}

type Component interface {
	search(string) []Component
}

type Department struct {
	components []Component
	name       string
}

func (d *Department) search(keyword string) []Component {
	//fmt.Println("Searching for keyword %s", keyword)
	var results []Component
	for _, composite := range d.components {
		res := composite.search(keyword)
		results = append(results, res...)
	}
	return results
}

func (d *Department) add(c Component) {
	d.components = append(d.components, c)
}

type Person struct {
	name string
}

func (p *Person) search(keyword string) []Component {
	//fmt.Printf("Searching for keyword %s in Person\n", keyword)
	if p.getName() == keyword {
		return []Component{p}
	}
	return nil
}

func (p *Person) getName() string {
	return p.name
}
