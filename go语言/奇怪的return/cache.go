package main

import "fmt"

//var cm = make(map[any]any)

type Cache[K comparable, V any] struct {
	M map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	c := &Cache[K, V]{
		M: make(map[K]V),
	}
	return c
}

func (c *Cache[K, V]) Set(k K, v V) {
	c.M[k] = v
}

func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	mv, ok := c.M[k]
	if ok {
		return mv, true
	} else {
		return
	}
}

func (c *Cache[K, V]) Get2(k K) (V, bool) {
	mv, ok := c.M[k]
	return mv, ok
}

func cache() {
	c := NewCache[string, int]()
	c.Set("one", 1)

	v, ok := c.Get("one")
	fmt.Println(v, ok)
	v, ok = c.Get("two")
	fmt.Println(v, ok)
	v, ok = c.Get2("two")
	fmt.Println(v, ok)
}
