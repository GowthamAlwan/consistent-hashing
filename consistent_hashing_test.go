package consistenthashing

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := consistentHash("hello")
	fmt.Println("hash = ", hash)
}
