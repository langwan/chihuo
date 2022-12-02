package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	machineIDBits      uint8 = 10
	sequenceNumberBits uint8 = 12
	epochTimestamp           = 1640995200 //2022-1-1 00:00:00
	machineMax         int64 = -1 ^ (-1 << machineIDBits)
	sequenceNumberMask int64 = -1 ^ (-1 << sequenceNumberBits)
)

type Snowflake struct {
	machineID      int64
	mu             sync.Mutex
	epoch          time.Time
	old            int64
	sequenceNumber int64
}

func (s *Snowflake) Gen() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Since(s.epoch).Milliseconds()
	if now == s.old {
		s.sequenceNumber = (s.sequenceNumber + 1) & sequenceNumberMask
		if s.sequenceNumber == 0 {
			for now <= s.old {
				now = time.Since(s.epoch).Milliseconds()
			}
		}
	} else {
		s.sequenceNumber = 0
	}
	s.old = now
	return now<<(machineIDBits+sequenceNumberBits) | (s.machineID << sequenceNumberBits) | s.sequenceNumber
}

func New(machineID int64) (*Snowflake, error) {

	if machineID < 0 || machineID > machineMax {
		return nil, errors.New("machine id must be between 0 and " + strconv.FormatInt(machineMax, 10))
	}

	s := Snowflake{machineID: machineID}
	s.epoch = time.Unix(epochTimestamp, 0)
	return &s, nil
}
