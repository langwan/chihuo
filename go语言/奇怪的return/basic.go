package main

import "fmt"

var m = make(map[string]int)

func get(k string) (int, bool) {
	if v, ok := m[k]; ok {
		return v, true
	} else {
		return v, false
	}
}

func get2(k string) (v int, ok bool) {
	v, ok = m[k]
	return
}

func basic() {
	m["one"] = 1

	val, ok := get("two")
	fmt.Println(val, ok)

	val, ok = get2("one")
	fmt.Println(val, ok)

	val, ok = get2("two")
	fmt.Println(val, ok)
}
