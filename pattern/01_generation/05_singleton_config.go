package main

import (
	"fmt"
)

/*
config
*/

type config struct {
	DBHost string
	DBPort int
	DBUser string
	DBName string
}

var configInstance *config

func getConfigInstance() *config {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now.")
				// 環境変数から取ってくる
				conf := config{
					"http://",
					5432,
					"user",
					"dbname",
				}
				configInstance = &conf
			})
	} else {
		fmt.Println("single instance already created.")
	}
	return configInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getOnceInstance()
	}
	fmt.Scanln()
}
