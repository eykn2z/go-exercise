// GOEXPERIMENT=range
package main

import (
	"fmt"
)

func main() {
	s := []string{"hello", "world"}
	for i, x := range slices.Backward(s) {
		fmt.Println(i, x)
	}
}
