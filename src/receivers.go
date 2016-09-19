package main

import "fmt"

type MyStruct struct {
	Field int
}

func (m MyStruct) ValueReceiver() {
	m.Field = 3
}

func (m *MyStruct) PointerReceiver() {
	m.Field = 4
}

func main() {
	m := MyStruct{
		Field: 1,
	}

	m.ValueReceiver()
	fmt.Println(m.Field)
	m.PointerReceiver()
	fmt.Println(m.Field)
}