package p2p

import (
	"fmt"
	"testing"
)

func TestStartRunner(t *testing.T) {
	StartRunner(10005)

	fmt.Printf("TestStartRunner: %s\n", "10001")
}
