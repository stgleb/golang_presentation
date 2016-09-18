package main

import "fmt"

func varArgs(args ...string) {
	for _, a := range args {
		fmt.Println(a)
	}
}

func main() {
	args := []string{"first", "second", "third"}
	varArgs(args...)
}
