package main

import (
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//LRU 建立有淘汰策略的缓存，一般指时间

// 单个key淘汰时间是1秒，5秒清洗一次
func TestBasic(t *testing.T) {

	c := cache.New(time.Second, 5*time.Second)

	c.Set("foo", "bar", cache.DefaultExpiration)
	foo, found := c.Get("foo")
	assert.True(t, found)
	assert.Equal(t, foo, "bar")

	time.Sleep(time.Second)
	foo, found = c.Get("foo")
	assert.False(t, found)
}

// 持久化
func TestNoExpiration(t *testing.T) {
	c := cache.New(time.Second, 5*time.Second)
	c.Set("foo", "bar", cache.NoExpiration)
	foo, found := c.Get("foo")

	assert.True(t, found)
	assert.Equal(t, foo, "bar")
	time.Sleep(time.Second)

	foo, found = c.Get("foo")
	assert.True(t, found)

	c.Delete("foo")
	foo, found = c.Get("foo")
	assert.False(t, found)
}

// 任何类型
func TestInterface(t *testing.T) {
	c := cache.New(time.Second, 5*time.Second)
	c.Set("foo", "bar", cache.NoExpiration)
	foo, found := c.Get("foo")
	assert.True(t, found)

	var str string
	str = foo.(string)
	assert.Equal(t, str, "bar")

	type Struct struct {
		Name string
	}
	c.Set("struct", &Struct{Name: "chihuo"}, cache.NoExpiration)

	var s *Struct
	foo, found = c.Get("struct")
	assert.True(t, found)
	s = foo.(*Struct)
	assert.Equal(t, s.Name, "chihuo")
}

// 指针造成线程安全的假象
func TestPointer(t *testing.T) {
	n := 10
	c := cache.New(time.Second, 5*time.Second)
	c.Set("foo", &n, cache.NoExpiration)
	foo, found := c.Get("foo")
	assert.True(t, found)
	assert.Equal(t, *(foo.(*int)), 10)
	n = 100
	foo, found = c.Get("foo")
	assert.True(t, found)
	assert.Equal(t, *(foo.(*int)), 100)
}

// 中断恢复
func TestNewFrom(t *testing.T) {
	store := map[string]cache.Item{"foo": {
		Object:     "bar",
		Expiration: int64(cache.NoExpiration),
	}}
	c := cache.NewFrom(time.Second, 5*time.Second, store)
	foo, find := c.Get("foo")
	assert.True(t, find)
	assert.Equal(t, foo, "bar")
}

// 自增
func TestIncrement(t *testing.T) {
	n := 0
	c := cache.New(time.Second, 5*time.Second)
	c.Set("foo", n, cache.NoExpiration)
	c.Increment("foo", 1)
	foo, find := c.Get("foo")
	assert.True(t, find)
	assert.Equal(t, foo, 1)

	c.Increment("foo", 2)
	foo, find = c.Get("foo")
	assert.True(t, find)
	assert.Equal(t, foo, 3)
}

// 遍历所有key
func TestItems(t *testing.T) {
	n := 0
	c := cache.New(time.Second, 5*time.Second)
	c.Set("foo", n, cache.NoExpiration)
	c.Set("foo2", n+1, cache.NoExpiration)
	items := c.Items()
	t.Log(items)
	assert.Equal(t, len(items), 2)
}
