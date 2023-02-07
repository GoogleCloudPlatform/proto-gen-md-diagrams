package proto

import (
	"flag"
	"testing"
)

func TestE2E(t *testing.T) {
	flag.Set("d", "data")
	flag.Set("w", "false")

	Execute()
}
