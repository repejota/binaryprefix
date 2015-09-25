package binaryprefix

import (
	"testing"
)

func TestParseBinaryPrefixMB(t *testing.T) {
	b, err := ParseBinaryPrefix("1MB")
	if err != nil {
		t.Error("Error parsing 1MB")
	}
	if int(b) != 1048576 {
		t.Error("1MB are 1048576 bytes but got ", int(b))
	}
}
