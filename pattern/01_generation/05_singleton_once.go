package main

import (
	"fmt"
	"sync"
)

/*
syncOnceパターン
*/

var once sync.Once

func getOnceInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("single instance already created.")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getOnceInstance()
	}
	fmt.Scanln()
}
