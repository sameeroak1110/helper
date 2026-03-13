package helper

import (
	"time"
	"math/rand"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) // seeding to nanoseconds granularity.

	if min >= max {
		return -1
	}

	//return rand.Intn(max - min) + min     // max is omitted. max limit is min through (max - 1.0)
	return rand.Intn(max - min + 1) + min   // max is honoured. max limit is min through max.
}
