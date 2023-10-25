package main

import (
	"fmt"
	"log"
)

/*
生成するプロダクトの各コンポーネントを抽象化し、一括生成する仕組み
大きい規模のプロダクトに使用

生き物等動くもので例えた方がいい？
*/

type Bed struct {
	Size     int
	Color    string
	Material string
}

type Desk struct {
	Width    int
	Depth    int
	Material string
}

type IBedFactory interface {
	createBed(size int) Bed
}

type IDeskFactory interface {
	createDesk() Desk
}

type IRoomFactory interface {
	createBedFactory() IBedFactory
	createDeskFactory() IDeskFactory
}

type ClassicBedFactory struct{}

func (c *ClassicBedFactory) createBed(size int) Bed {
	return Bed{Size: size, Color: "white", Material: "wood"}
}

type ClassicDeskFactory struct{}

func (c *ClassicDeskFactory) createDesk() Desk {
	return Desk{Width: 100, Depth: 60, Material: "wood"}
}

type ClassicRoomFactory struct {
	name string
}

func (c *ClassicRoomFactory) createBedFactory() IBedFactory {
	return &ClassicBedFactory{}
}
func (c *ClassicRoomFactory) createDeskFactory() IDeskFactory {
	return &ClassicDeskFactory{}
}

func getRoomFactory(roomType string) (IRoomFactory, error) {
	if roomType == "classic" {
		return &ClassicRoomFactory{name: "classic room"}, nil
	}
	return nil, fmt.Errorf("invalid room type: %s", roomType)
}

func main() {
	roomType := "classic"
	factory, err := getRoomFactory(roomType)
	if err != nil {
		log.Fatal(err)
	}
	bedFactory := factory.createBedFactory()
	bed := bedFactory.createBed(120)
	deskFactory := factory.createDeskFactory()
	desk := deskFactory.createDesk()
	log.Printf("roomType: %s, bed: %v, desk: %v", roomType, bed, desk)
}
