package goroutine

import (
	"log"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var sharedData int = 0

	go func() {
		sharedData++
	}()

	if sharedData == 0 {
		log.Printf("Data is %v", sharedData)
	}
}
