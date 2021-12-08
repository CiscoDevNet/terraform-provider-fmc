package fmc

import (
	"fmt"
	"math/rand"
	"time"
)

// generates random string with fixed size
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
