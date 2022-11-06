package main

import (
	"app/core"
	"fmt"
)

// go run -ldflags "-X app/core.Version=1.0.1 -X app/core.Build=20221106" .
// go build -ldflags "-X app/core.Version=1.0.1 -X app/core.Build=20221106" .
func main() {
	fmt.Printf("version = %s, build = %s\n", core.Version, core.Build)
}
