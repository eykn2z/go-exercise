package main

import "fmt"

//shape.areaの例
//https://gist.github.com/c-yan/451a59fed7448cdec79eb0ef33fcc9a7

type Painter interface {
	Paint()
	String() string
}

type Writer interface {
	Write()
	String() string
}

func displayArt(p Painter) {
	fmt.Printf("Painter can display art: %s\n", p)
}

func publishWriting(w Writer) {
	fmt.Printf("Writer can publish writing: %s\n", w)
}

type Artist struct {
	name string
}

func (a Artist) Paint() {
	fmt.Println("Artist is painting a picture")
}

func (a Artist) Write() {
	fmt.Println("Artist is writing a book")
}

func (a Artist) String() string {
	return a.name
}

func main() {
	artist := Artist{"John"}

	displayArt(artist)
	publishWriting(artist)
}
