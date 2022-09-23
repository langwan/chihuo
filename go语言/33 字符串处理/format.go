package main

import "fmt"

type Value struct {
	Name  string
	Value int32
}

func f1() {
	v1 := Value{
		Name:  "val",
		Value: 10,
	}
	s1 := fmt.Sprintf("%d %v %+v %#v %T %p %f", v1.Value, v1, v1, v1, v1, &v1, float64(v1.Value))
	fmt.Printf("format is %s\n", s1)
}
