package main

import (
	"testing"
	"time"
)

func TestTimeSub(t *testing.T) {
	start := time.Now()
	end := time.Now()
	elapsed := end.Sub(start)
	t.Log(elapsed)
}

func TestTimeSince(t *testing.T) {
	start := time.Now()
	t.Log(time.Since(start))
}

func TestTimeEcho(t *testing.T) {
	t.Log(time.Now().String())
	time.Sleep(200 * time.Millisecond)
	t.Log(time.Now().String())
}
