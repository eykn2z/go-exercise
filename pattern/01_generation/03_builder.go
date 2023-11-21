package main

import "fmt"

/*
構成物：
builder,director
builder：任意のパーツ制作の置き場
director：同パーツを組み合わせて生成するもの: クライアントコードからプロダクト構築を隠蔽する
directorの中にパーツの組み込みを行うので段階的に構築が可能

class1(ex: House)が持つclass member(class2,3..)(ex: 1st floor, 2nd floor...)が複数ある時、その設定をclass1のconstructorでやってしまうと複雑になる
そのため HouseBuilderを作成し、build1stFloor(),build2ndFloor()の中で引数を決め、
その任意の引数を持ってHouseを作成するメソッドをdirectorで管理する

※各builderのメソッドの引数に固定値を入れられれば（パターン化出来れば)、builderは実装可能
※あくまで複数のbuilderパターンが存在する・増えうる時に実装にしないと、複雑性が増す

abstract factory = builderのbuilder?
*/

type IBuilder interface {
	//setGraphicBoard()
	//setCPU()
	//setMemory()
	//setHDD()
	//setPCCase()
	getPC() PC
}

func getBuilder(builderType string) IBuilder {
	if builderType == "gaming" {
		return newGamingPCBuilder()
	}

	if builderType == "cheap" {
		return newCheaperPCBuilder()
	}
	return nil
}

type PC struct {
	graphicBoardType string
	cpuType          string
	memoryType       string
	hddType          string
	pcCaseType       string
}

type CheaperPCBuilder struct{}

func newCheaperPCBuilder() *CheaperPCBuilder {
	return &CheaperPCBuilder{}
}

func (b *CheaperPCBuilder) getPC() PC {
	return PC{
		graphicBoardType: "",
		cpuType:          "intel core 2 duo",
		memoryType:       "ddr2 8GB",
		hddType:          "HDD 64GB",
		pcCaseType:       "cheapest black",
	}
}

type GamingPCBuilder struct{}

func newGamingPCBuilder() *GamingPCBuilder {
	return &GamingPCBuilder{}
}

func (b *GamingPCBuilder) getPC() PC {
	return PC{
		graphicBoardType: "geforce 7800",
		cpuType:          "intel core i9",
		memoryType:       "ddr3 32GB",
		hddType:          "SDD 256GB",
		pcCaseType:       "gaming black",
	}
}

type Director struct {
	builder IBuilder
}

func newPCDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}
func (d *Director) buildPCs(times int) []PC {
	var pcs []PC
	for i := 0; i < times; i++ {
		pcs = append(pcs, d.builder.getPC())
	}
	return pcs
}

func main() {
	gamingPCBuilder := getBuilder("gaming")
	cheaperPCBuilder := getBuilder("cheap")
	director := newPCDirector(gamingPCBuilder)
	gamingPCs := director.buildPCs(10)

	fmt.Printf("Gaming PC builder: %s\n", len(gamingPCs))

	director = newPCDirector(cheaperPCBuilder)
	cheaperPCs := director.buildPCs(20)

	fmt.Printf("Cheaper PC builder: %s\n", len(cheaperPCs))
}
