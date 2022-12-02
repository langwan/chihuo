package main

import (
	"fmt"
	"testing"
	"time"
)

const (
	machineIDBits      uint8 = 10
	sequenceNumberBits uint8 = 12
	epoch              int64 = 1640995200 //2022-1-1 00:00:00
)

func TestMock(t *testing.T) {

	epochTime := time.Unix(epoch, 0)

	fmt.Printf("%64b int64 max\n", 9223372036854775807)
	var timestamp int64 = time.Since(epochTime).Milliseconds()
	fmt.Printf("%64b timestamp\n", timestamp)
	var machineID int64 = 1
	fmt.Printf("%64b machineID\n", machineID)
	var sequenceNumber int64 = 1
	fmt.Printf("%64b sequenceNumber\n", sequenceNumber)

	fmt.Printf("\n")

	fmt.Printf("%64b sequenceNumber\n", sequenceNumber)
	timestamp = timestamp << (machineIDBits + sequenceNumberBits)
	fmt.Printf("%64b timestamp shift\n", timestamp)
	machineID = machineID << sequenceNumberBits
	fmt.Printf("%64b machineID shift\n", machineID)
	fmt.Printf("%64b sequenceNumber\n", sequenceNumber)
	id := int64(timestamp | machineID | sequenceNumber)
	fmt.Printf("%64b id\n", id)
}
