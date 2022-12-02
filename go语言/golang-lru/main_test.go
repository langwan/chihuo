package golang_lru

import (
	glcache "github.com/hashicorp/golang-lru/v2"
	"testing"
)

func Test1(t *testing.T) {
	cache, err := glcache.New[int, string](128)
	if err != nil {
		t.Error(err)
		return
	}
	ok := cache.Add(1, "langwan")
	if ok == true {
		t.Fail()
		return
	}
	value, ok := cache.Get(1)
	if ok == false {
		t.Fail()
		return
	}
	t.Log(value)
}
