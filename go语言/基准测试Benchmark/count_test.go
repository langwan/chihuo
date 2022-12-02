package main

/**
并行基准测试
*/
import (
	"sync"
	"sync/atomic"
	"testing"
)

func Benchmark_MutexCount(b *testing.B) {
	var count uint64 = 0
	var lock sync.Mutex
	count = 0
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			count++
			lock.Unlock()
		}
	})
	b.StopTimer()
}

func Benchmark_AtomicCount(b *testing.B) {

	var count uint64 = 0
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint64(&count, 1)
		}
	})
	b.StopTimer()

}

func Benchmark_RWLockCount(b *testing.B) {

	var count uint64 = 0
	var lock sync.RWMutex

	count = 0
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			count++
			lock.Unlock()
		}
	})
	b.StopTimer()
}
