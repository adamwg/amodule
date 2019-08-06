package blah

import "fmt"

type Foo struct {
	name string
}

func NewFoo(name string) *Foo {
	return &Foo{name}
}

func (f *Foo) Dance(times int) {
	for i := 0; i < times; i++ {
		fmt.Println(f.name)
	}
}
