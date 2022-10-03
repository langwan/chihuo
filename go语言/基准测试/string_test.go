package main

/**
基准测试默认的执行时间是1秒
串行基准测试
*/
import (
	"strings"
	"testing"
)

func Benchmark_StringCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := ""
		for i := 0; i < 25565; i++ {
			str += "a"
		}
	}
	//	b.Log(b.N)
}

func Benchmark_BufferCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for i := 0; i < 25565; i++ {
			b.WriteString("a")
		}
		b.String()
	}
	b.Log(b.N)
}
