package main

import "fmt"

type Foo struct {
	Value int64
}

func (f *Foo) Incr() {
	f.Value++
}

func (f Foo) Incr2() {
	f.Value++
}

func (Foo) TableName() string {
	return "tb_foo"
}

type Tabler interface {
	TableName() string
}

func main() {
	foo := Foo{Value: 0}

	foo.Incr()
	fmt.Printf("value is %d\n", foo.Value)

	foo.Incr2()
	fmt.Printf("value is %d\n", foo.Value)

	if tabler, ok := interface{}(foo).(Tabler); ok {
		fmt.Printf("value is %s\n", tabler.TableName())
	} else {
		fmt.Printf("value is %s\n", "foo")
	}
}
