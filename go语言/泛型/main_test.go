package main

import (
	"fmt"
	"golang.org/x/exp/maps"
	"testing"
)

func printT[T any](s []T) { // Just an example, not the suggested syntax.
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func printI(s []interface{}) { // Just an example, not the suggested syntax.
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func Test1(t *testing.T) {
	printT[int]([]int{1, 2, 3, 4, 5})
	printT[string]([]string{"a", "b", "c", "d", "e"})
}

func Test2(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	keys := maps.Keys(m)
	printT[int](keys)
}

func Test3(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	keys := maps.Keys(m)
	printT[int](keys)
}

type Cache[K comparable, V any] struct {
	M map[K]V
}

func (c *Cache[K, V]) Set(k K, v V) {
	m[k] = v
}

func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	mv, ok := m[k]
	if ok {
		return mv.(V), true
	} else {
		return
	}
}

func Test4(t *testing.T) {
	c := Cache[string, int]{}
	c.Set("one", 1)
	c.Set("two", 2)
	v, ok := c.Get("one")
	t.Log(v, ok)
	v, ok = c.Get("three")
	t.Log(v, ok)
}
