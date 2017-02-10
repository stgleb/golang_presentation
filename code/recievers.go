// https://play.golang.org/p/zHINE9lpk0

package main

import (
	"fmt"
)

type Foo struct{
	value int
}


func (f Foo) bar() string {
	f.value = 3
	return fmt.Sprintf("bar from %v", f)
}

func (f *Foo) baz() string {
	f.value = 4
	return fmt.Sprintf("baz from %v", f)
}

func main() {
	f := Foo{1}
	f2 := &Foo{2}
	fmt.Println(f.bar())
	fmt.Println(f)
	fmt.Println(f.baz())
	fmt.Println(f)

	fmt.Println(f2.bar())
	fmt.Println(f2)
	fmt.Println(f2.baz())
	fmt.Println(f2)
}
