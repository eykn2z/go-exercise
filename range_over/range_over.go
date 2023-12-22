// GOEXPERIMENT=range
package main

import (
	"fmt"
)

func main() {
	s := []string{"hello", "world"}
	for i, x := range s {
		fmt.Println(i, x)
	}
}
