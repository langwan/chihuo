package main

import "fmt"

var cm = make(map[any]any)

type Cache[K comparable, V any] struct {
	M map[K]V
}

func (c *Cache[K, V]) Set(k K, v V) {
	cm[k] = v
}

func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	mv, ok := cm[k]
	if ok {
		return mv.(V), true
	} else {
		return
	}
}

func (c *Cache[K, V]) Get2(k K) (V, bool) {
	var v V
	mv, ok := cm[k]
	if ok {
		return mv.(V), true
	} else {
		return v, false
	}
}

func cache() {
	c := Cache[string, int]{}
	c.Set("one", 1)

	v, ok := c.Get("one")
	fmt.Println(v, ok)
	v, ok = c.Get("two")
	fmt.Println(v, ok)
	v, ok = c.Get2("two")
	fmt.Println(v, ok)
}
