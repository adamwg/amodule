package blah

import "fmt"

// Foo is a struct.
type Foo struct {
	name string
}

// NewFoo returns a Foo.
func NewFoo(name string) *Foo {
	return &Foo{name}
}

// Dance monkey dance.
func (f *Foo) Dance(times int) {
	for i := 0; i < times; i++ {
		fmt.Println(f.name)
	}
}
