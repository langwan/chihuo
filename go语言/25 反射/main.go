package main

import (
	"fmt"
	"reflect"
)

type Service struct {
	Name string `json:"name"`
}

func (s Service) GetName() string {
	return s.Name
}

func (s *Service) SetName(name string) {
	s.Name = name
}

func worker1() {
	s := Service{}
	s.SetName("chihuo")
	name := s.GetName()
	fmt.Printf("call GetName return %s\n", name)
}

func worker2() {
	s := Service{}
	rv := reflect.ValueOf(&s)
	params := []reflect.Value{reflect.ValueOf("chihuo")}
	rv.MethodByName("SetName").Call(params)

	ret := rv.MethodByName("GetName").Call(nil)

	fmt.Printf("reflect call return %s\n", ret[0].String())
}

func worker3() {
	s := Service{}
	rt := reflect.TypeOf(s)
	if field, ok := rt.FieldByName("Name"); ok {
		tag := field.Tag.Get("json")
		fmt.Printf("field tag is %s\n", tag)
	}
}

func main() {
	worker1()
	worker2()
	worker3()
}
