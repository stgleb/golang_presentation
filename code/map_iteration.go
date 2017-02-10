// https://play.golang.org/p/LILMXTTEMn

package main

import (
	"fmt"
)

func main() {
	m := map[int]interface{}{
		1: nil,
		2: nil,
		3: nil,
	}

	for i := 0;i < 5; i++ {
		for key := range m {
			fmt.Println(key)
		}
		fmt.Println("------------------")
	}
}
